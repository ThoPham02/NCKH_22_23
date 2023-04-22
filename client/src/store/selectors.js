// select login info
export const loginSelector = (state) => state.login;
export const userSelector = (state) => state.login.user;
export const roleSelector = (state) => state.login.user.typeAccount;
export const tokenSelector = (state) => state.login.token;

