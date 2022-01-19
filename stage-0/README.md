# Software Architecture

When starting any project it is important to understand what you need to do and what tools you need to do it. Planning is important but you don't necessarily always need a flow chart or diagram to really decide.

Most developers  in school learn of only one language (that's how it should be to avoid confusion). However, a lot of the times in software development, more than one language is used.

 In the software world, using many different tools that are really good at doing a couple things as opposed to using one tool thats decent at everything has proven to be better.

![Sample Software Architecture Image by Litslink](https://litslink.com/wp-content/uploads/2021/04/Web_Application_Architecture_Diagram__diagram_.png)
This example is much more complex than ours is going to be because our video chat app will simply be one server with no external persistent data.

The model of this project looks more like this:
```
  [user] [user] [user] [user]  ... n number of users
    ^     ^       ^      ^
    |     |       |      |
    |     |       |      |
    v     v       v      v
  [----------server---------]
```
In this model, the users in their browser will send chat messages or video/audio data and the server will merely echo this back to everyone.

To break this system down into parts as the model abstracts a lot of the details, all apps can be broken down into two parts: a frontend and backend.

## Frontend
A frontend is essentially the main thing people think of when they thing of any app. The frontend is the design of an app and it is how the code is presented to someone using the system. It is important to decide this first because it doesn't matter if your code is fast and correct if nobody can use it. 

![Command Line Application (CLI)](https://miro.medium.com/max/2718/1*4jGCY6YznCuRlYiLPaL27A.png)

![Mobile Frontend](https://cdn.dribbble.com/users/1615584/screenshots/14038579/media/852e5fd6b106616e00e78c25870041d1.jpg?compress=1&resize=400x300)

![Android Studio Layout Designer](https://i.stack.imgur.com/exGNG.png)

![Swift UI Designer](https://docs-assets.developer.apple.com/published/a151730046f7ac186031a760fe890b92/11800/overview-hero@2x.png)

![Web/Desktop Frontend](https://i.pinimg.com/originals/b3/44/aa/b344aa836c6536b9324bbc4e449e0697.jpg)

![C# for Windows Apps on Visual Studio Community](https://i.stack.imgur.com/P7xTw.jpg)

Many different frameworks exist to try and streamline the design process and provide accessibility across all platforms. The browser is limited to HTML/CSS/JS while desktop and mobile are limited to their native platforms.

Web frameworks:
- ReactJS
- AngularJS
- VueJS

Universal:
- Flutter

In our project we will be keeping it simple and using plain HTML/CS/JS for our frontend.

## Backend
The backend of an application is the actual functionality of an app. The backend takes requests from a frontend whether it be pressing a button and then getting data from the server, starting a download, etc.

In our case, the backend will be doing most of the work handling requests from the frontend to join the call, taking in the video/audio feed, broadcasting video along with chat messages.

For simplicity, we're going to be using the programming language Go or Golang because it simplifies a lot of the process of creating a backend. Most languages require you to use a 3rd party API but Go has included a web server package in the standard library.

# [Web Tech Stacks](https://youtu.be/Sxxw3qtb3_g)

Tech stacks are essentially the system of different technologies used to provide a product.

### Basic Example:
Imagine an online shop made by a non-developer:
- WiX / Squarespace / Wordpress (Frontend)
- PayPal / Shopify (Backend)

Most popular examples of stacks:
- **LAMP** (Linux, Apache, MySQL, PHP/Perl/Python)
- **MEAN** (MongoDB, Express, Angular, NodeJS)
- **MERN** (MongoDB, Express, React, NodeJS)
- **MEVN** (MongoDB, Express, Vue, NodeJS)

# Resources for our Tech Stack
*You don't need to watch the videos in their entirety, they are primarily here as a reference in case you want to look something up or actually master these technologies.* 

*For this project, a basic understanding of these technologies is enough*
- [Learn HTML/CSS Video](https://youtu.be/mU6anWqZJcc)
- [Learn JavaScript Video](https://youtu.be/PkZNo7MFNFg)
- [W3Schools](https://www.w3schools.com/)
- [Learn Golang Video](https://youtu.be/YS4e4q9oBaU)
- [Go Examples](https://gobyexample.com/)
- [Go Documentation](https://pkg.go.dev/std)