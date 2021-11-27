import { AxiosInstance } from "axios";
import coreIns from "../instances/core";

type IPersonRepository = {
  redirectGoogle: (
    redirect_url: string
  ) => Promise<{ oauth_url: string; state: string }>;
};

export const PersonRepository = (axios: AxiosInstance): IPersonRepository => ({
  redirectGoogle: async (redirect_url) =>
    (await axios.post("/auth/login/google", { redirect_url })).data,
});

export default PersonRepository(coreIns);
