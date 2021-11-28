import Card from "../../component/cards/rounded";
import Password from "../../component/input/password";
import NoLayout from "../../component/layout/noLayout";
import PageProps from "../../definition/pageProps";
import BtnGoogle from "../../component/button/google";
import { useState } from "react";
import Username from "../../component/input/username";
import styles from "./login.module.scss";
import { cat } from "../../util/utils";
import { Primary as BtnPrimary } from "../../component/button/mui";
import { motion } from "framer-motion";

const Page: PageProps = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = () => {};

  return (
    <div className={styles.background}>
      <div className={styles.loginpage}>
        <motion.div
          animate={{
            scale: [0.8, 1.1, 1],
            opacity: [0, 0.8, 1],
          }}
          transition={{
            duration: 1,
            ease: "easeInOut",
            times: [0, 0.5, 1],
            repeat: 0,
          }}
        >
          <Card.Simple
            className={cat(styles.card, "col-11 col-sm-8 col-md-7 col-lg-5")}
          >
            <h1 className={cat(styles.titlebar, "mx-auto")}>
              Login to <span className={styles.gotitle}>GoChat</span>
            </h1>
            <Username
              errmsg={password.length === 0 ? "password must be filled" : ""}
              className="mt-1 w-100"
              label="username"
              value={username}
              onChange={setUsername}
            />
            <Password
              errmsg={password.length === 0 ? "password must be filled" : ""}
              className="mt-3 w-100"
              label="password"
              value={password}
              onChange={setPassword}
            />
            <div className="mt-3 mx-auto w-50">
              <BtnPrimary onClick={handleLogin} className="w-100">
                login
              </BtnPrimary>
            </div>
            <p className="mx-auto my-3 w-fit">{`< OR />`}</p>
            <div className="mt-3 mx-auto w-fit">
              <BtnGoogle />
            </div>
          </Card.Simple>
        </motion.div>
      </div>
    </div>
  );
};

Page.layout = NoLayout;
export default Page;
