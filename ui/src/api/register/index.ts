import {post} from "/@/utils/request";

const baseUrl = "/api"
export const registerApi = (p: object) => post(baseUrl + "/register", p);
