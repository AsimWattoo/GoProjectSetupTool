// @ts-nocheck
import express from "express";
require('dotenv').config();
import 'tsconfig-paths/register';
import cors from 'cors';
import bodyParser from 'body-parser';
import cookieParser from "cookie-parser";
import session from "express-session";
import path from "path";
import baseRoute from "@routes/index";
import RequestLogger from "@middlewares/request-logger.ts";
import RateLimiter from "@middlewares/rate-limiter.ts";
import ErrorHandler from "@middlewares/error-handler.ts";
import Logger from "@services/Logger.ts";
import config from '@/settings/config';

// Creating the app
const app = express();

// Setting Up Public Folder
const rootDir = path.resolve(__dirname, '../')
app.use('/uploads', express.static(path.join(rootDir, 'uploads')));

//Setting Up Middlewares
app.use(bodyParser.json({limit: "100mb"}))
app.use(bodyParser.urlencoded({ extended: true }));

// Setting Up Cookie Parsing
app.use(cookieParser());

// Setting Up CORS
Logger.LogImportant(`Setting CORS Origins to ${config.ORIGINS.join(', ')}`)
app.use(cors({origin: [], credentials: true}))

// Setting Up Sessions
const store = new session.MemoryStore();
app.use(session({
    secret: "test123",
    resave: false,
    saveUninitialized: true,
    store: store,
    cookie: {
        maxAge: 60 * 60 * 3,
        httpOnly: true,
        secure: true,
        path: '/',
        sameSite: 'none'
    }
}))

// Setting Up Routes
app.use('/api', RequestLogger(), RateLimiter(), baseRoute);

// Adding an error handler
app.use(ErrorHandler());

app.listen(config.PORT, () => {
    Logger.LogMessage(`Server listening at http://127.0.0.1:${config.PORT}`);
})

export default app;