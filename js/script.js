// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

function calculatePressed() {
  //This function converts Fahrenheit into Celsius
  //Input through Textfields
  const fahrenheit = parseFloat(document.getElementById("fahrenheit").value)

  //Process
  const celsius = (fahrenheit - 32) * 5/9


  //Output
  document.getElementById("answer").innerHTML =
    "This Fahrenheit temperature in Celsius is: " + celsius.toFixed(2) + "°C"
}
