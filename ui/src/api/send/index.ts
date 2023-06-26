import {post} from "/@/utils/request";

const baseUrl = "/owl"
export const sendMailApi = (p: object) => post(baseUrl + "/api/v2/messages/mail", p);
