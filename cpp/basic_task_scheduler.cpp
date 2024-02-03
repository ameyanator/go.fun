#include <iostream>
#include <vector>
#include <functional>
#include <thread>
#include <condition_variable>
#include <queue>

using namespace std;

class BasicTaskScheduler {
public:
    BasicTaskScheduler() {
        worker = std::thread([this] { workerThread(); });
    }

    ~BasicTaskScheduler() {
        stopWorker = true;
        cond.notify_one();
        worker.join();
    }

    void schedule(std::function<void()> task) {
        std::unique_lock<std::mutex> locker(mu);
        tasks.push(task);
        locker.unlock();

        cond.notify_one();
    }

    void waitUntilComplete() {
        std::unique_lock<std::mutex> locker(mu);
        cond.wait(locker, [this]{return tasks.empty();});
    }

private:

    void workerThread() {
        while(!stopWorker) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return stopWorker || !tasks.empty();});
            if (stopWorker) {
                return;
            }
            auto task = tasks.front();
            tasks.pop();
            locker.unlock();

            task();

            locker.lock();
            if(tasks.empty()) {
                cond.notify_one();
            }
        }
    }
    std::condition_variable cond;
    queue<std::function<void()> > tasks;
    std::thread worker;
    bool stopWorker = false;
    std::mutex mu;
};

// Example usage
int main() {
    BasicTaskScheduler taskScheduler;

    // Schedule tasks
    for (int i = 1; i <= 5; ++i) {
        taskScheduler.schedule([i]() {
            std::cout << "Task " << i << " started." << std::endl;
            // Simulate some work
            std::this_thread::sleep_for(std::chrono::seconds(1));
            std::cout << "Task " << i << " completed." << std::endl;
        });
    }
    cout<<"waiting for tasks now"<<endl;
    // Wait until all tasks are completed
    taskScheduler.waitUntilComplete();

    std::cout << "All tasks completed." << std::endl;

    return 0;
}