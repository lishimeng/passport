import {get, post} from "/@/utils/request";

const baseUrl = "/api"
export const signInApi = (p: object) => post(baseUrl + "/sign_in/login", p);
export const signInCodeApi = (p: object) => post(baseUrl + "/sign_in/codeLogin", p);
export const getUserInfoApi = (p: object) => get(baseUrl + "/user/getUserInfo", p);
export const checkSignInApi = (p: object) => post(baseUrl + "/sign_in/checkToken", p);