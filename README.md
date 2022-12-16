# ETI Assignment 1 (Master)
Name: Lee Wen Kang<br />
Class: P03<br />
ID: 10203100B<br />

## Contents

1. [Repositories](#Repositories)
2. [Features and Design Considerations](#Features-and-Design-Considerations)
3. [Solution Architecture](#Solution-Architecture)
4. [Trip Assignment Process](#Trip-Assignment-Process)


This assignment is to implement a ride-share platform using a microservice architecture with 2 primary group of users, passangers and drivers. Passangers should be able to start trips while drivers should be able to accept them.

## Repositories
| No        | Service Name           | Purpose  | Link  |
| :------------- |:-------------| :-----| :-----|
| 1 | rideshare-cli (current) | Acts as an interface for users to interact with. It connects to rideshare-api to interact with the database. | [Link](https://github.com/NPLeeWenKang/rideshare-cli) |
| 2 | rideshare-api | Interacts directly with the database for persistant data storage. Uses REST. | [Link](https://github.com/NPLeeWenKang/rideshare_api_svc) |
| 3 | rideshare-tripassignment | Service that is in charge of assigning trips to drivers. | [Link](https://github.com/NPLeeWenKang/rideshare-tripassignment) |
| 4 | rideshare-db | MySQL for persistant data storage. | [Link](https://github.com/NPLeeWenKang/rideshare_db) |
| 5 | rideshare-ui (bonus) | For the bonus marks, this service serves a website built using React. | [Link](https://github.com/NPLeeWenKang/rideshare-ui) |

## Features and Design Considerations

## Solution Architecture

Before any code has been written, the entity relations and the overall architecture was drawn out to easily understand and scale the project. Furthermore, planning early reducing the need to refactor large chunks of code whenever new requirements are discovered.

### Entity Relationship Diagram
<img src="https://user-images.githubusercontent.com/73012553/208162989-6a729f6d-0611-40fd-9365-fcd159d1ef5f.png" alt="Architecture Diagram" width="800"/>

For the RideShare project, there are a total of 4 entities, Passanger, Trip, Driver and Trip Assignment. The requirements for the entity attributes have been gathered from the assignment brief. 

However, for the Trip Assignment entity, I took liberty in coming up with the attributes needed to satisfy the design considerations stated before. As seen, there is a seperation of relationship between Trip and Driver via Trip Assignment as this would allow drivers to reject trip assignments without affecting the Trip entity. By seperating this, it also normalises the data.

### Architecture Diagram
<img src="https://user-images.githubusercontent.com/73012553/208163133-07261890-11ba-493c-8da6-7772240ea376.png" alt="Architecture Diagram" width="500"/>

Because the project adopted a microservice architecture, several services has been created.

* **rideshare-cli** - Built with GO, this service acts as an interface for users to interact with the RideShare system. It has the appropriate error checks and satisfies all the functionalities listed above.

* **rideshare-api** - Built with GO, this service interacts with RideShare's database and allows other services to communicate with it via REST. This service is live on port 5000.

* **rideshare-tripassignment** - Built with GO, this service is in charge of handling the trip<>driver assignments where it runs the assignment algorithem every 8 seconds. It is good to take note that this service does not have any exposed ports and connects directly with the database instead of via the api.

* **rideshare-db** - For persistant data storage, a MySQL database was used. Although not required by the assignment, this service has been configured to run on Docker enviroments. Because MySQL's default port is 3306, this has been kept the same with Docker's exposed port being set to 3306:3306.

* **rideshare-ui (bonus)** - A web interface has been created with React that allows users to interact with the RideShare via their referred browser instead of a CLI. The web UI mimics the CLI interface with identical control flow, display style and functionalities. Because this service is a "bonus", this service has been developed in and only tested on Chrome Version 108.0.5359.125 (Official Build) (64-bit).

## Trip Assignment Process

While designing this website, I had to think of my users and how they would use and navigate the website. As a result, I tried to imagine and understand my users from their point of view.
Firstly, I identified some potential users that would browse my website, and why they would want to use my website.
1. As a scholarship interviewer, I would like to know more about the interviewee. These can include his past projects, a short description about himself or his thinking mindset. By knowing more  about the interviewee, I would be able to better judge if he will make use of the scholarship to it's fullest extent.
2. As a University , I would like to understand more about the student. I would like to know his past education institutes, some of his notable achievements and his past projects. This allows me to determine if he is passionate about the course and whether to accept his university application. <br/>

To ensure that the website was mobile friendly, I had to design and develop the website from a "Mobile First" perspective. This means that the mobile view is always the first priority. This ensures that the website is always mobile and PC friendly.
When I was drawing my wireframe using Adobe Xd, I had to ensure that both PC and mobile views where suitable and appropriate.<br/>

As a web developer, I have a personal "rule" I always follow. That "rule" is to always develop bit size code and always check and ensure that they contain no errors. By developing in this manner, it is easier to troubleshoot and error check, as it is easier to look and detect errors in a small chunk of code than a large chunk of code. This makes developing any software whether is it a website or a programme much easier and faster.
![development process](/github-README-src/development-process.PNG?raw=true)

Developing a wireframe is vital, as it help me plan out the website layout and navigation. This gives me an overview of my website, and allows me to "follow" the wireframe layout. This makes development much more efficient and quick.

* PC version of XD wireframe: https://xd.adobe.com/view/2e1a46c6-e03c-4f04-9d58-0ca583770d8d-7b49/

* Mobile version of XD wireframe: https://xd.adobe.com/view/a5842d5c-794d-4e71-b890-e5b6ada51c64-f0d7/

## Features
Since this is a personal portfolio website. I have implemented several features to ensure that the navigation is smooth and experience is great.<br/>
1. Viewing personal project / source code of project. In each personal project page, there is a link either to access the project or the project's Github page. One example is in the "Todo Master" project page. In that page I provided a link (https://todo-app-bb61a.firebaseapp.com/) that brings the user a place where the user can "experience" the project.
2. Another feature is to view each technology for each project. This allows users to click on the icons which brings them to the technology's branding page. One example is in the "Box With You" project page. At the bottom, it showcases the technologies used in that project. By clicking on the icon, the user will be brought to the technology's branding page. By clicking on the Firebase icon, the user will be brought to https://firebase.google.com/.<br/>
## Testing
Testing is a vital part in progamming and website development. It ensures that what we are building works and is functional. Therefore, comprehensive testing methods has to be in place to help expose errors.<br/>
### Online validators
To check and test my website, I used w3school's CSS and Markup service. This allowed me to validate my code and ensure that there are no errors.
1. W3C CSS Validation Service: https://jigsaw.w3.org/css-validator/
2. W3C Markup Validation Service: https://validator.w3.org/
### Screen size
As the website is being hosted on the browser, both mobile and PC users would be able to access it. So to ensure that navigation and viewabiity is good on all screens, steps has to be taken to ensure that everything works. I have detail several methods I have used to test different screen sizes.<br/>
1. Resizing browser manually<br/>
By resizing the browser manually, I am able to quickly view my website in different view heights and widths. Although this may not be a full proof way to test screen sizes, it allows for quick and easy testing (testing while developing).
2. Using Chrome's mobile emulator<br/>
By inspecting the page, a button will appear. This button toggles Chrome's mobile emulator, allowing you to view the screen through different phones. This is a quick test to ensure that the contents in the screen are readable and properly aligned. However this method has several limitations, as it does not actually run the website on a mobile phone, but only simulates it. To read more about Chrome's mobile emulator: https://developers.google.com/web/tools/chrome-devtools/device-mode.
3. Using a physical phone<br/>
Apart from using the other methods to test my website in different screen size, I also used my Samsung J7 phone to view the website. This allows me to accurately experience how navigating and viewing the contents on a mobile phone feels like.
### Browsers
As HTML and CSS may be compiled differently in different browsers, it is vital test the website on different browsers. As I developed my website on Chrome, I went and tested my website on FireFox and Microsoft Edge. This allowed me to ensure that my website was functional on different browsers.
### Links
To test my links, I had to ensure that internal links did not open a new page, while external links like social media and Github links all opened a new page when clicked on. This is vital, as it improves the user experience.
### Contact form
To test the contact form, I had to perform several steps.
1. Open the "Contact Me" page and click submit. A message should appear saying "Please fill up this form" under the email imput section.
2. Input an invalid email, for example "user.com". A message should appear saying that the email has to contain a "@".
3. After properly filling the email section, click submit again. Another message will appear asking the user to fill up the "Content" section.
4. After filling the form with the proper email and content, click submit. After clicking it, the default system email application will appear, showing the message. The reciepent of the email should be "spaceycodes@gmail.com".
## Credits
### Icons
* Bootstrap Icons (Arrow): https://icons.getbootstrap.com/icons/arrow-left-short/
* Bootstrap Icons (Envelope): https://icons.getbootstrap.com/icons/envelope/
* Bootstrap Icons (Link): https://icons.getbootstrap.com/icons/link-45deg/
* Ngee Ann Logo: https://www.np.edu.sg/Pages/default.aspx
* Christ Church Secondary logo: https://christchurchsec.moe.edu.sg/our-school/motto-n-crest
* Si Ling Primary logo: https://silingpri.moe.edu.sg/
* Github logo: https://github.com/logos
* LinkedIn logo: https://brand.linkedin.com/downloads
* React logo: https://reactjs.org/
* React-beautiful-dnd: https://github.com/atlassian/react-beautiful-dnd
* Firebase logo: https://firebase.google.com/brand-guidelines
* Javascript logo: https://github.com/voodootikigod/logo.js/
* HTML 5 logo: https://www.w3.org/html/logo/
* Flutter logo: https://flutter.dev/brand
* Dart logo: https://dart.dev/brand
### Guides
* w3schools (Timeline): https://www.w3schools.com/howto/howto_css_timeline.asp
### Certain parts of the website was inspired by
* Nadia Campo Woytuk: https://nadiacw.com/
