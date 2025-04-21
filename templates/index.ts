// @ts-nocheck
import express from "express";
require('dotenv').config();
import 'tsconfig-paths/register';
import cors from 'cors';
import bodyParser from 'body-parser';
import cookieParser from "cookie-parser";
import session from "express-session";
import path from "path";

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

app.get('/test', (req, res) => {
    res.status(200).json({data: "This is the test route"})
})

const PORT = 3000;

app.listen(PORT, () => {
    console.log(`Server listening at http://127.0.0.1:${PORT}`);
})

export default app;