import Cookies from "js-cookie"
const tokenKey = "systemToken"
export const getToken = () => Cookies.get(tokenKey)
export const setToken = (token) => Cookies.set(tokenKey, token)
export const removeToken = () => Cookies.remove(tokenKey)
