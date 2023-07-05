



[//]: # (<!-- PROJECT SHIELDS -->)

[//]: # (<!--)

[//]: # (*** I'm using markdown "reference style" links for readability.)

[//]: # (*** Reference links are enclosed in brackets [ ] instead of parentheses &#40; &#41;.)

[//]: # (*** See the bottom of this document for the declaration of the reference variables)

[//]: # (*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.)

[//]: # (*** https://www.markdownguide.org/basic-syntax/#reference-style-links)

[//]: # (-->)

[//]: # ([![Contributors][contributors-shield]][contributors-url])

[//]: # ([![Forks][forks-shield]][forks-url])

[//]: # ([![Stargazers][stars-shield]][stars-url])

[//]: # ([![Issues][issues-shield]][issues-url])

[//]: # ([![MIT License][license-shield]][license-url])

[//]: # ([![LinkedIn][linkedin-shield]][linkedin-url])



<!-- PROJECT LOGO -->
<br />
<div align="center">
    <a href="https://github.com/Tobiska/logity"><img width="400" height="60" src="assets/logo.svg" alt="logity" border="0"></a>
  <p align="center">
    <br />

[//]: # (    <a href="https://github.com/github_username/repo_name"><strong>Explore the docs »</strong></a>)

[//]: # (    <br />)

[//]: # (    <br />)

[//]: # (    <a href="https://github.com/github_username/repo_name">View Demo</a>)

[//]: # (    ·)

[//]: # (    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>)

[//]: # (    ·)

[//]: # (    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a>)
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#environment-logity">Environment logity</a></li>
        <li><a href="#centrifugo">Centrifugo</a></li>
         <li><a href="#neo4j">Neo4j</a></li>
        <li><a href="#run">Run</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>

[//]: # (    <li><a href="#acknowledgments">Acknowledgments</a></li>)
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Backend service for minimalistic social network logity. A user in logity communicates with others using chat rooms, each user can send a log (analogue of a message in social networks). Each log can contain a photo, text or picture.

### Feature
- Authentication / Authorization User with using **JWT** tokens and **Postgresql**.
- Maintain relationships between users and manage their chat rooms **Neo4j**.
- Real-time messaging management using **Websocket** and **Centrifugo**.

### Built With

[![Go][Go]][Go-url]
[![JWT][JWT]][JWT-url]
[![Postgres][Postgres]][Postgres-url]
[![Neo4j][Neo4j]][Neo4j-url]

<!-- GETTING STARTED -->
## Getting Started

### Environment logity

1. Create file **dev.env** from **template.env** and fill variables with your values
   - `YOUR_POSTGRES_USERNAME`, `YOUR_POSTGRES_PASSWORD` - postgres credentials, which you specify in environment postgres **docker** container.
   - `YOUR_LOGITY_USERNAME`, `YOUR_LOGITY_PASSWORD` - specify user credentials for neo4j-database **logity**.
   - `YOUR_CENTRIFUGO_API_KEY` - the api key by which **logity** to access **centrifugo**.
   - `YOUR_CENTRIFUGO_SECRET_KEY` - the secret key which **logity** will sign the client's **JWT** tokens for access to the **centrifugo**.
   - `YOUR_SECRET_ACCESS_KEY` - the secret key which **logity** will sign access tokens.
   - `YOUR_SECRET_REFRESH_KEY` - the secret key which **logity** will sign refresh tokens.
2. To run locally without container, we recommended create another one file **local.env**.

### Centrifugo

1. Create file **config.json** from **exports/template_config.json** and fill variables with your values
   - `YOUR_CENTRIFUGO_SECRET_KEY` - the secret key which centrifugo will verify client's JWT tokens. Must be equal logity environment variable `YOUR_CENTRIFUGO_SECRET_KEY`
   - `YOUR_CENTRIFUGO_API_KEY` - the api key which centrifugo will verify service's JWT token. Must be equal logity environment variable `YOUR_CENTRIFUGO_API_KEY`.
   - `YOUR_ADMIN_PASSWORD`, `YOUR_ADMIN_SECRET` - admin credentials.
2. Property ```token_issuer``` must be equal env variable `APP_NAME`.

### Neo4j

1. Run the docker-container with neo4j: ```docker-compose up neo4j --build```
2. Go to http://localhost:7474/browser/.
3. Authenticate with **neo4j / testify**. 
4. In command line ``` CREATE DATABASE logity;``` and ```:use logity```
5. Create user with role **architect** using ```:server user add``` and check it using ```:server user list```

### Run

1. Run the docker-containers: ```docker-compose up --build```
2. Neo4j UI: http://localhost:7474/browser/.
3. Centrifugo admin panel: http://localhost:9123/.
4. If liquibase containers fail, try running them later. If successful, then postgres or neo4j did not have time to start to run.

<!-- USAGE EXAMPLES -->
## Usage

### REST API
 **logity** has a **REST** API.
1. OpenApi docs: http://localhost:8080/swagger/index.html.
2. The **postman** v2.1 collection is also stored in the **scripts/postman** folder.

### Log Client
To test the operation of real-time messages, you can use a simple go-client that uses the **centrifugo** API.

1. Go to `scripts/client/client.go`
2. SignIn by route `http://<<**host**>>/auth/sign-in` and copy
3. From response copy `rtc_token` into `const Token = <TOKEN>`
4. Run client.go
5. Update subscribes using `http://<<**host**>>/op/update-subscribes`

<!-- ROADMAP -->
## Roadmap

- [ ] Transactions Manager. Need to manage transactions from many sources(neo4j postgres).
- [ ] Role model.
- [ ] Logging, Tracing.

[//]: # (See the [open issues]&#40;https://github.com/github_username/repo_name/issues&#41; for a full list of proposed features &#40;and known issues&#41;.)


[//]: # (<!-- LICENSE -->)

[//]: # (## License)

[//]: # ()
[//]: # (Distributed under the MIT License. See `LICENSE.txt` for more information.)

[//]: # ()
[//]: # (<p align="right">&#40;<a href="#readme-top">back to top</a>&#41;</p>)



<!-- CONTACT -->
## Contact

Your Name - [@Tobiska](https://t.me/Tobiska) - tobiskaKirill@gmail.com

Project Link: [https://github.com/Tobiska/logity](https://github.com/Tobiska/logity)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Neo4j]: https://img.shields.io/badge/Neo4j-008CC1?style=for-the-badge&logo=neo4j&logoColor=white
[Neo4j-url]: https://neo4j.com/

[Postgres]: https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white
[Postgres-url]: https://www.postgresql.org/

[JWT]: https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens
[JWT-url]: https://jwt.io/

[Go]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/

