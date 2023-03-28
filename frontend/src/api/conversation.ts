import { request } from "@playwright/test";
import Api from "../services/Api";

export type ConvoMessage = {
    messageId: number;
    message: string;
    sentByMe: boolean;
}

export const ConverseWithDad = async (requestMessages: ConvoMessage, userId: number ) => {
    try {
        let request = {
            "request": requestMessages.message,
            "user_id": userId,
        }
        return Api.post("/conversation", request);
    } catch (error) {
        console.error(error);
        return error
    }
}
