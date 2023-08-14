import {get, post} from "/@/utils/request";

const baseUrl = "/api"
export const signInApi = (p: object) => post(baseUrl + "/sign_in/login", p);
export const signInCodeApi = (p: object) => post(baseUrl + "/sign_in/codeLogin", p);
export const getUserInfoApi = (p: object) => get(baseUrl + "/user/getUserInfo", p);
export const checkSignInApi = (p: object) => post(baseUrl + "/sign_in/checkToken", p);
export const clearTokenApi = (p: object) => post(baseUrl + "/sign_in/clearToken", p);

export const bindSendCodeApi = (p: object) => get(baseUrl + '/user/bindSendCode', p)
export const bindPhoneApi = (p: object) => post(baseUrl + "/user/bindPhone", p);
export const bindEmailApi = (p: object) => post(baseUrl + "/user/bindEmail", p);
export const changePasswordApi = (p: object) => post(baseUrl + "/user/changePassword", p);
