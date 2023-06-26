import {post} from "/@/utils/request";

const baseUrl = "/api"
export const signInApi = (p: object) => post(baseUrl + "/sign_in", p);
export const signInCodeApi = (p: object) => post(baseUrl + "/sign_in/code", p);