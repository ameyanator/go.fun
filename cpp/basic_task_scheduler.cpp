#include <iostream>
#include <functional>
#include <thread>
#include <queue>
#include <mutex>
#include <condition_variable>
#include <ctime>
#include <vector>

using namespace std;

class Task {
public:
    int id;
    std::function<void()> func;
    std::time_t execution_time;

    void print()
    {
        std::cout << "boo" << '\n';
    }
};

class BasicScheduler {
public:

    BasicScheduler() {
        worker = std::thread(poller);
    }

    ~BasicScheduler() {
        worker.join();
    }

    void schedule(Task task) {
        std::unique_lock<std::mutex> locker(mu);
        heap.push(task);
        locker.unlock();
        cond.notify_all();
    }
private:

    void poller() {
        while(true) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return !heap.empty();});
            Task task = heap.top(); heap.pop();
            locker.unlock();
            task.func();
        }
    }

std::thread worker;
std::mutex mu;
std::condition_variable cond;
priority_queue<Task, vector<Task>, greater<> > heap;
};

class TimeBasedScheduler {
    TimeBasedScheduler() {
        worker = std::thread(poller);
    }

    ~TimeBasedScheduler() {
        worker.join();
    }

    void schedule(Task task) {
        std::unique_lock<std::mutex> locker(mu);
        heap.push(task);
    }

private:

    void poller() {
        while(true) {
            const auto now = std::chrono::system_clock::now();
            const std::time_t t_c = std::chrono::system_clock::to_time_t(now);
            std::unique_lock<std::mutex> locker(mu);
            if(heap.top().execution_time < t_c) {
                locker.unlock();
                std::this_thread::sleep_for(std::chrono::milliseconds(1000));
                continue;
            }

            Task task = heap.top(); heap.pop();
            locker.unlock();
            task.func();
        }
    }
std::thread worker;
std::mutex mu;
std::condition_variable cond;
priority_queue<Task, vector<Task>, greater<> > heap;
};

class TopoScheduler {
public:

    TopoScheduler() {
        worker = std::thread(poller);
        worker1 = std::thread(getNewTasks);
        new_tasks_added = true;
    }

    ~TopoScheduler() {

    }

private:

    /**
     * [] -> [1,2,3]
     * [1] -> [5,7]
     * [1,2,7] -> [8,9]
     * 
    */

   int counter = 0;

    vector<Task> NewTasks(vector<Task> tasks) {
        return {new Task()};
    }

    void getNewTasks() {
        while(true) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return new_tasks_added;});
            vector<Task> new_tasks = NewTasks(already_executed);
            new_tasks_added = false;
            locker.unlock();

            locker.lock();
            for(auto t : new_tasks) {
                to_be_executed.push(t);
            }
            locker.unlock();
            cond.notify_all();
        }
    }

    void poller() {
        while(true) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return !to_be_executed.empty();});
            Task task = to_be_executed.front(); to_be_executed.pop();
            locker.unlock();
            task.func();

            locker.lock();
            already_executed.push_back(task);
            new_tasks_added = true;
            locker.unlock();
            cond.notify_all();
        }
    }

    void getNewTasks1() {
        while(true) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return new_tasks_added;});

            vector<Task> new_tasks = NewTasks(already_executed);
            new_tasks_added = false;

            // Lock mu before accessing to_be_executed
            locker.lock();
            for(auto t : new_tasks) {
                to_be_executed.push(t);
            }
            locker.unlock();

            cond.notify_all();
        }
    }

    

    

    void poller1() {
        while(true) {
            std::unique_lock<std::mutex> locker(mu);
            cond.wait(locker, [this](){return !to_be_executed.empty();});
            Task task = to_be_executed.front(); to_be_executed.pop();
            locker.unlock();
            task.func();

            // Lock mu before accessing already_executed
            locker.lock();
            already_executed.push_back(task);
            new_tasks_added = true;
            locker.unlock();

            cond.notify_all();
        }
    }


std::thread worker;
std::mutex mu;
std::condition_variable cond;
queue<Task> to_be_executed;

std::thread worker1;
vector<Task> already_executed;
bool new_tasks_added;
};