import { AxiosInstance } from "axios";
import coreIns from "../instances/core";

type IPersonRepository = {
  redirectGoogle: (
    redirect_url: string
  ) => Promise<{ oauth_url: string; state: string }>;
};

export const PersonRepository = (axios: AxiosInstance): IPersonRepository => ({
  redirectGoogle: async (redirect_url) => {
    try {
      return (await axios.post("/auth/login/google", { redirect_url })).data;
    } catch (err) {
      console.error(err);
      return { oauth_url: "", state: "" };
    }
  },
});

export default PersonRepository(coreIns);
