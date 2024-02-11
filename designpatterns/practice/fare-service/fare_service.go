package main

type FareService interface {
	getCost(*Trip) float64
}