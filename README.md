# MoviesListAPI
**Go** + **Gin** + **Gorm** Backend API for handling user registration/login (using JWT for auth) and Review CRUD functions.

Movie IDs that correspond to those found in The Movie Database's API.

<br >
  
# 📁 Collection: User 

<br >

## End-point: RegisterUser
Registers a user.
#### Method: 🟨 POST 🟨
>```
>http://localhost:8080/register
>```
#### Body (**json**)

```json
{
    "username": "test2",
    "password": "test"
}
```  

<br >

## End-point: GetUserByID
Returns the user with corresponding user ID.
#### Method: 🟩 GET 🟩
>```
>http://localhost:8080/user/6
>```

<br >
  
## End-point: UpdateUser
Updates specified user's login details. (WIP).
#### Method: 🟦 PUT 🟦
>```
>http://localhost:8080/user/4
>```
#### Body (**json**)

```json
{
    "username": "test2",
    "password": "test"
}
```

<br >

## End-point: DeleteUser
Deletes a user with the corresponding user ID.
#### Method: 🟥 DELETE 🟥
>```
>http://localhost:8080/user/3
>```

<br >
  
## End-point: LoginUser
Logs the user in.

If login attempt is successful, a cookie containing a JWT called "jwt" is returned in the response.
#### Method: 🟨 POST 🟨
>```
>http://localhost:8080/login
>```
#### Body (**json**)

```json
{
    "username": "test",
    "password": "test"
}
```

<br >
  
# 📁 Collection: Review 

<br >

## End-point: CreateReview
Generates a review.
#### Method: 🟨 POST 🟨
>```
>http://localhost:8080/review
>```
#### Body (**json**)

```json
{
    "authorId": 2,
    "movieId": "19995",
    "rating": 2.5,
    "subject": "i love this film!!",
    "body": "FUNNY blue men!"
}
```

<br >
  
## End-point: GetReviewsByMovieID
Returns all reviews with specified movie ID.
#### Method: 🟩 GET 🟩
>```
>http://localhost:8080/review/19995
>```

<br >
  
## End-point: DeleteReview
Deletes review with corresponding review ID.
#### Method: 🟥 DELETE 🟥
>```
>http://localhost:8080/review/5
>```

<br >

## End-point: UpdateReview
Updates review with corresponding review ID.

Custom middleware prevents user from editing reviews they didn't author.
#### Method: 🟦 PUT 🟦
>```
>http://localhost:8080/review/2
>```
#### Body (**json**)

```json
{
    "movieId": "19995",
    "rating": 4.0,
    "subject": "i LOVE this film!!",
    "body": "Funny blue men!"
}
```

<br >
  
## End-point: GetAvgRatingByMovieID
Returns the average rating of reviews containing specified Movie ID as a float.

\-1 is returned if no reviews have been created with the specified Movie ID.
#### Method: 🟩 GET 🟩
>```
>http://localhost:8080/review/19995/rating
>```

<br >
  
## End-point: GetMoviesByAuthor
Returns a list of all of the reviews with corresponding username.
#### Method: 🟩 GET 🟩
>```
>http://localhost:8080/review/user/username
>```
