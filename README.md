# Grat-Project-POC

using for prove of concept for Computer Engineer Project Preparation

## frontend

This directory is a part of our website frontend written in React TS.
library using
- yarn 
- vite
- axios
- tailwind

to run *must be in frontend directory*
```
yarn dev
```

may require 
```
yarn 
```

## backend

This directory is a part of our website backend written in go. using go echo framework.

to run *must be in backend directory*
```
go run ./cmd/main
```

## model

This directory is part of our AI model written in python. using fastapi as an api for communicate with backend.

to run *must be in model directory*
```
uvicorn main:app --reload
```

may require
```
pip install fastapi
pip install "uvicorn[standard]"
```