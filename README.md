# funStoreHub Golang API Authentication

Welcome to the funStoreHub Golang API Authentication documentation. This API powers the backend for your favorite club's sports center store, allowing seamless interaction with the funStoreHub frontend. It enables users to create, display, edit, and delete products specific to their beloved team.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Authentication Endpoints](#authentication-endpoints)
- [Protected Endpoints](#protected-endpoints)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This section briefly introduces the funStoreHub Golang API Authentication.

## Features


- User authentication using JWT (JSON Web Tokens)
- Secure token generation and validation
- Endpoints for user login and authentication

## Installation

```bash
go get github.com/gin-gonic/gin
go get github.com/pilu/fresh
go get golang.org/x/crypto
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
```