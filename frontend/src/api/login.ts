import Api from "../services/Api";

export const Login = async (username: string, password: string) => {
    try {
        let request = {
            "username": username,
            "password": password
        }
        // const response = await Api.post("/login", request);
        return Api.post("/login", request)
    } catch (error) {
        console.error(error);
        return "Unsuccesful Login Attempt"
    }
};
