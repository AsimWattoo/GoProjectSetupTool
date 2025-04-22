// @ts-nocheck
import moment from "moment";
import DateFormats from "@/data/DateFormats.ts";

type LogMessage = string | object | boolean | number

class Logger {

    static LogImportant(message: LogMessage) {
        const date = new Date();
        console.warn(`[${moment(date).format(DateFormats.Log)}] => ${message}`)
    }

    static LogMessage(message: LogMessage) {
        const date = new Date();
        console.log(`[${moment(date).format(DateFormats.Log)}] => ${message}`)
    }

    static LogError(message: LogMessage) {
        const date = new Date();
        console.error(`[${moment(date).format(DateFormats.Log)}] => ${message}`)
    }
}

export default Logger;