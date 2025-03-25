# SimpleWall — A Microservices Platform for Content Sharing and Interaction
🇷🇺 [Русская версия]([README.md](https://github.com/talgatbaytleu/SimpleWall/blob/docs/rus/README_RUS.md))

SimpleWall is a backend platform that allows users to share photos, interact with content, and track activity.

## Key Features:
- User registration and post publishing with descriptions
- Editing and deleting own posts
- Viewing all posts
- Liking and unliking posts
- Displaying the number of likes and list of likers
- Commenting on posts and sending notifications to the frontend
- Editing and deleting own comments
- Viewing a specific comment or a list of comments for a post
- Retrieving a list of all user posts with a count of likes and comments

The project is based on a microservices architecture and does not include a frontend. It provides an HTTP API for data interaction.

## Project Structure
The application consists of 7 microservices, each serving a specific function:

![SimpleWall-diagram](https://github.com/user-attachments/assets/6d823b5a-e715-4c29-8136-3ff4d1c79368)

- **Gateway** — Request routing (monolithic architecture).
- **Auth** — User authentication with JWT tokens (three-tier architecture).
- **Post** — Post management (three-tier architecture).
- **Like** — Processing likes and unlikes (three-tier architecture).
- **Comment** — Comment management (three-tier architecture).
- **Wall** — News feed assembly (hexagonal architecture).
- **Notification** — Notifications (three-tier architecture).

### Endpoints

***auth-service:***
- **POST /register** – Register a new user.
- **POST /login** – Login to receive a token.
- **POST /validate** – Validate a token.

***post-service:***
- **POST /post** – Publish a post.
- **GET /post/{post_id}** – Retrieve a post by `post_id`.
- **DELETE /post/{post_id}** – Delete a post by `post_id` (if you are the owner).
- **PUT /post/{post_id}** – Edit a post by `post_id` (if you are the owner).

***like-service:***
- **POST /like** – Like a post, `post_id` is specified in the request body.
- **GET /likes/count?post_id={post_id}** – Retrieve the number of likes for `post_id`.
- **GET /likes?post_id={post_id}** – Retrieve the list of likers for `post_id`.
- **DELETE /like** – Unlike a post, `post_id` is specified in the request body.

***comment-service:***
- **POST /comment** – Write a comment for a post, `post_id` is specified in the request body.
- **PUT /comment/{comment_id}** – Edit a comment by `comment_id` (if you are the owner).
- **DELETE /comment/{comment_id}** – Delete a comment by `comment_id` (if you are the owner).
- **GET /comment/{comment_id}** – Retrieve a comment by `comment_id`.
- **GET /comments?post_id={post_id}** – Retrieve a list of comments for `post_id`.

***wall-service:***
- **GET /wall?user_id={user_id}** – Retrieve the user's post feed by `user_id`.

### Technologies Used
- **PostgreSQL** — Two databases are used:
  - `sw_users_auth` for authentication.
  - `sw_posts_db` for storing posts, likes, and comments.

![ERD_for_SW](https://github.com/user-attachments/assets/1a3a3ade-e438-482f-9f0b-aa546c6fcf43)

- **Redis** — Caching for fast feed loading (Wall service).
- **Kafka** — Message transmission (`comment-kafka-notification`).
- **S3 Prototype** — A custom (poor but functional) implementation of a binary file storage system.
- **Docker** — All services are deployed in containers.
- **Git** — Development was done through separate branches, simulating real team collaboration.
