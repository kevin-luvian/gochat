import axios from "axios";
import { store } from "../../redux/store";
import { selectAccessToken } from "../../redux/states/auth";

const makeCoreInstance = () => {
  const instance = axios.create({
    baseURL: "http://localhost:8000/",
  });

  instance.interceptors.request.use(async (config) => {
    if (config.headers !== undefined)
      config.headers.Authorization = selectAccessToken(store.getState());
    return config;
  });

  return instance;
};

export default makeCoreInstance();
