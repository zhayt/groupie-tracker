# About project
***
## groupie-tracker
### Objectives

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

    - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

    - The second one, `locations`, consists in their last and/or upcoming concert locations.

    - The third one, `dates`, consists in their last and/or upcoming concert dates.

    - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.

- This project also focuses on the creation of events/actions and on their visualization.

    - The event/action we want you to do is known as a client call to the server (client-server). We can say it is a feature of your choice that needs to trigger an action. This action must communicate with the server in order to recieve information, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)
    - An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.


## How to run
If you are in the root directory of the current project ~/../groupie-tracker, then:
```shell
$ go run cmd/* -addr=:9000
INFO    2023/02/20 04:26:35 Starting server on http://localhost:9000
...
```
Follow the link that output in terminal, like this [http://localhost:9000](http://localhost:9000)
