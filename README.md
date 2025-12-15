# Groupie Tracker

## 📌 Project Overview

The **Groupie Tracker** project is about building a **user-friendly website** that displays information about artists and bands by consuming a provided API.
It combines **data visualization**, **client-server communication**, and **event-driven programming** to create an interactive experience.

The API contains four main parts:

1. **Artists** – Information about bands/artists: name, image, start year, first album date, and members.
2. **Locations** – Last and/or upcoming concert locations.
3. **Dates** – Last and/or upcoming concert dates.
4. **Relation** – Links between artists, locations, and dates.

Using this data, the goal is to build a website where the information can be displayed in a **clear and interactive way**, using cards, tables, lists, graphics.

---

## 🎯 Objectives

The main objectives of this project are:

* Consume and manipulate a **RESTful API**.
* Build a **dynamic website** displaying artists, concerts, and relationships between them.
* Create **events/actions** that trigger communication with the server (client-server).
* Ensure a **robust backend** that does not crash under normal usage.
* Implement **good coding practices** in Go.
---

## ⚙️ Instructions

* The **backend must be written in Go**.
* Handle all errors gracefully; the site and server must **never crash**.
* Display the data in a way that is **user-friendly and intuitive**.
* Events/actions should be implemented to demonstrate **client-server communication**, such as fetching or updating data upon a user action.
* You are allowed to use **only standard Go packages**.
* Ensure **all pages work correctly** and maintain a consistent structure.

---

## 💻 Allowed Packages

* Only **standard Go packages** are permitted.
* No external libraries or frameworks for Go are allowed.

---

## 📂 Usage

* The site fetches data from the provided API and displays it dynamically.

* Example of RESTful API structure:

  * `/artists` – returns artist information
  * `/locations` – returns concert locations
  * `/dates` – returns concert dates
  * `/relation` – links data from the above endpoints

* The frontend displays the data using **HTML and visual elements**.

* Events on the page (like clicking a card or button) should **trigger requests** to the backend to fetch or manipulate data.

---

## 🚧 Difficulties I Faced

During this project, I encountered several challenges:

* **Data relationships**
  Understanding how artists, locations, dates, and relations are connected required careful mapping.

* **Event creation**
  Implementing client-server communication for triggered actions (requests/responses) was tricky at first.

* **Data visualization**
  Designing the UI to display large amounts of data in a clear, readable way was challenging.

* **Error handling**
  Ensuring the server never crashes, even with unexpected inputs, was a constant focus.

* **Go standard library only**
  Building a complete backend without external libraries required careful planning.

---

## 📚 What I Learned

This project helped me strengthen my knowledge in several areas:

* **Manipulating JSON data** from an API in Go
* **Storing and linking data** between multiple sources
* Building **dynamic HTML pages** to display complex data
* Creating **client-server events**
* Writing **robust and maintainable Go code**
* Handling **errors gracefully**
* Basics of **UI/UX for data display**

---

## ✅ Conclusion

The **Groupie Tracker** project taught me how to combine **backend Go programming** with **frontend data visualization**, while respecting coding standards and robust error handling.
It was a great opportunity to understand **client-server communication** and how to create **interactive web applications** from real API data.

