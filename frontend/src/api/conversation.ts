import { request } from "@playwright/test";
import Api from "../services/Api";

export interface ConvoMessage {
    messageId: number;
    messages: ConvoStruct[]
}

export interface ConvoStruct {
    message: string;
    role: string;
    messageID: number;
}

// export CreateConvoMessage = (Messages: ConvoMessage[], newMessage: string) => {}

export const ConverseWithDad = async (requestMessages: ConvoMessage, userId: number, ) => {
    try {
        if (requestMessages.messages.length > 0) {
        let request = {
            "request": requestMessages?.messages?.at(-1)?.message,
            "user_id": userId,
            "messages": requestMessages?.messages
        }
        console.log(request);
        return Api.post("/conversation", request);
    }
    } catch (error) {
        console.error(error);
        return error
    }
}
