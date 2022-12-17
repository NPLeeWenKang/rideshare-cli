# ETI Assignment 1 (Master)

Name: Lee Wen Kang<br />
Class: P03<br />
ID: 10203100B<br />

## Contents

1. [Repositories](#Repositories)
2. [Requirements and Design Considerations](#Requirements-and-Design-Considerations)
3. [Solution Architecture](#Solution-Architecture)
4. [Startup Guide](#Startup-Guide)
4. [Proof of Completion](#Proof-of-Completion)

This assignment is to implement a ride-share platform using a microservice architecture with 2 primary group of users, passangers and drivers. Passangers should be able to start trips while drivers should be able to accept them.

## Repositories

---

| No        | Service Name           | Purpose  | Link  |
| :------------- |:-------------| :-----| :-----|
| 1 | rideshare-cli (current) | Acts as an interface for users to interact with. It connects to rideshare-api to interact with the database. | [Link](https://github.com/NPLeeWenKang/rideshare-cli) |
| 2 | rideshare-account-svc | Interacts directly with the database for persistant data storage for passangers and drivers. Uses REST. | [Link](https://github.com/NPLeeWenKang/rideshare-account-svc) |
| 3 | rideshare-trip-svc | Interacts directly with the database for persistant data storage for trips and its assignments. Uses REST. | [Link](https://github.com/NPLeeWenKang/rideshare-trip-svc) |
| 4 | rideshare-ta_process-svc | Service that is in charge of assigning trips to drivers. Trip assignment is abbreviated as ta.| [Link](https://github.com/NPLeeWenKang/rideshare-ta_process-svc) |
| 5 | rideshare-system-db | MySQL for persistant data storage. | [Link](https://github.com/NPLeeWenKang/rideshare-system-db) |
| 6 | rideshare-ui (bonus) | For the bonus marks, this service serves a website built using React. | [Link](https://github.com/NPLeeWenKang/rideshare-ui) |

## Requirements and Design Considerations

---

Quote from assignment brief

> You are required to implement a ride-sharing platform using microservice architecture. The platform has 2 primary group of users, namely the passengers and drivers. Users can create either account. 
>
> During creation of passenger account, first name, last name, mobile number, and email address are required. Subsequently, users can update any information in their account, but they are not able to delete their accounts for audit purposes.
>
> For driver account creation, first name, last name, mobile number, email address, identification number and car license number are required. Drivers can update all information except their identification number. Similarly, a driver account cannot be deleted.
>
> A passenger can request for a trip with the postal codes of the pick-up and drop-off location. The platform will assign an available driver, who is not driving a passenger, to the trip. This driver will then be able to initiate a start trip or end trip. The passenger can retrieve all trips he/she has taken before in reverse chronological order

### Requirements

---

Having analysed the assignment brief, these are some of the requirements gathered and will be implemented.

1. **Select user to "login"** - As an authentication system is needed, the system will simply request the user to input the user Id to "login" as.

2. **Create passanger** - Allows users to create passanger entities using the attributes, first name, last name, mobile number and email address. Passanger Ids are to be auto assigned.

3. **Create driver** - Allows users to create driver entities using the attributes, first name, last name, mobile number, email address, identification number and car number. Driver Ids are to be auto assigned.

4. **Update passanger** - Allows users to edit all passanger information.

5. **Update driver** - Allows users to edit all driver information except the identification number.

6. **Passanger and driver cannot be deleted** - For auditing purposes, users cannot delete any entities.

7. **Display trips for passanger** - Display the trips taken by a passanger in descending order based on the trip id.

8. **Create/start trip** - Passangers should be able to start a trip by specifying their pick-up and drop-off location.

9. **Trip assignment** - The system should assign unassigned trips to drivers (in-depth explaination in the [Design Considerations](#Design-Considerations) section).

10. **Display current trip assignments (passanger)** - Passangers should be able to see all their current trips that are currently still in progress (in-depth explaination in the [Design Considerations](#Design-Considerations) section). They should be able to see:

    * Trip id.
    * Driver's id, first name, last name and mobile number.
    * Trip's pick-up location, drop-off location, start time, end time and status.

11. **Display current trip assignments (driver)** - Driver should be able to see their current trips that are currently still in progress (in-depth explaination in the [Design Considerations](#Design-Considerations) section). They should be able to see:

    * Trip id
    * Passanger's id, first name, last name and mobile number.
    * Trip's pick-up location, drop-off location, start time, end time and status.

12. **Driver should be able to change status of his/her trip** - This can include rejecting, accepting, starting and ending trips. At each status of the trip, the RideShare system should handle it appropriately [Design Considerations](#Design-Considerations) section).

13. **(Bonus) Website UI** - A web UI to replace the CLI interface.

### Design Considerations

---

### Passanger can create multiple trips

To mimic actual ride-share platforms, the system has been created to allow passangers to create new trips while their current trip is still in progress, allowing passangers to have more flexibility in booking rides when they require multiple cars.

To add on, although passangers can start multiple trips simultaneously, drivers can **ONLY** have one trip assignment at once. 

### Definition of "in progress" trips

To ensure that the display of "in progress" trips is correct, it has been defined as:

* Trip must have a status of `PENDING`, `ACCEPTED` or `DRIVING`.

### Drivers are able to reject trip assignments

To facilitate a better driver experience, drivers are able to reject the trip assigned to them. At this point, the trip assignment system will then look for another "available" drivers (definition of "available" specified below).

### Definition of "available" drivers

To ensure that the trip assignment of drivers to trips is reliable, the definition of "available" has been defined as:

* Driver must be available. Driver must have set their `is_available` attribute to `true`.
* Driver must not be occupied with a "in progress" trip.
* Driver must not have rejected the trip before.

### Different trip statuses

* **Pending** - Trip has been assigned to the driver but the driver has not accepted it yet.
* **Rejected** - Driver has rejected the trip assignment and the trip assignment algorithem should reassign another available driver.
* **Accepted** - Trip has been assigned to the driver and the driver has accepted the assignment. However, the driver has yet to pick-up the passanger
* **Driving** - Driver has picked up the passanger and is currently driving to the drop-off location.
* **Done** - Driver has arrived at the drop-off location and the trip is finished.

### Trip assignment process

![image](https://user-images.githubusercontent.com/73012553/208237499-7c74c161-cbed-4864-9003-b914c24ba988.png)

![image](https://user-images.githubusercontent.com/73012553/208237503-dee4a213-b508-4148-8005-c5c3226c962e.png)

The trip assignment process for this system is quite different. Instead of assigning trips to drivers at the point of trip creation or rejection, the service runs a trip assignment process every 8 seconds.

It first gets all the unassigned trips, then assign "available" drivers to the trips.

If the trip assignment is rejected by the driver, the trip is than placed back into the unassigned pool of trips and are ready to be assigned by the service again.

## Solution Architecture

---

Before any code has been written, the entity relations and the overall architecture was drawn out to easily understand and scale the project. Furthermore, planning early reducing the need to refactor large chunks of code whenever new requirements are discovered.

### Entity Relationship Diagram

<img src="https://user-images.githubusercontent.com/73012553/208162989-6a729f6d-0611-40fd-9365-fcd159d1ef5f.png" alt="Architecture Diagram" width="1000"/>

For the RideShare project, there are a total of 4 entities, Passanger, Trip, Driver and Trip Assignment. The requirements for the entity attributes have been gathered from the assignment brief. 

However, for the Trip Assignment entity, I took liberty in coming up with the attributes needed to satisfy the design considerations stated before. As seen, there is a seperation of relationship between Trip and Driver via Trip Assignment as this would allow drivers to reject trip assignments without affecting the Trip entity. By seperating this, it also normalises the data.

### Architecture Diagram

<img src="https://user-images.githubusercontent.com/73012553/208241443-594c1790-28f8-47e1-be51-f88ba8609dd9.png" alt="Architecture Diagram" width="700"/>

Because the project adopted a microservice architecture, several services has been created.

* **rideshare-cli** - Built with GO, this service acts as an interface for users to interact with the RideShare system. It has the appropriate error checks and satisfies all the functionalities listed above.

* **rideshare-account-svc** - Built with GO, this service interacts with RideShare's database and allows other services to communicate with it via REST. This service is in charge of accounts such as passangers and drivers. This service is live on port 5000.

* **rideshare-account-svc** - Built with GO, this service interacts with RideShare's database and allows other services to communicate with it via REST. This service is in charge of trips and its assignments. This service is live on port 5000.

* **rideshare-ta_process-svc** - Built with GO, this service is in charge of handling the trip<>driver assignments where it runs the assignment algorithem every 8 seconds. Take note that this service does not have any exposed ports and connects directly with the database instead of via other services.

* **rideshare-system-db** - For persistant data storage, a MySQL database was used. Although not required by the assignment, this service has been configured to run on Docker enviroments. Because MySQL's default port is 3306, this has been kept the same with Docker's exposed port being set to 3306:3306.

* **rideshare-ui (bonus)** - A web interface has been created with React that allows users to interact with the RideShare via their referred browser instead of a CLI. The web UI mimics the CLI interface with identical control flow, display style and functionalities. Because this service is a "bonus", this service has been developed in and only tested on Chrome Version 108.0.5359.125 (Official Build) (64-bit).

## Startup Guide

---

## Proof of Completion

---
