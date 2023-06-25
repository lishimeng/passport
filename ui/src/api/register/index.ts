import {post} from "/@/utils/request";

const baseUrl = "/api"
export const registerApi = (p: object) => post(baseUrl + "/register", p);
export const signInApi = (p: object) => post(baseUrl + "/sign_in", p);