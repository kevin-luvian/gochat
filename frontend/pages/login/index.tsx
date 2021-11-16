import Card from "../../component/cards/rounded";
import Password from "../../component/input/password";
import OneLayout from "../../component/layout/oneLayout";
import PageProps from "../../definition/pageProps";
import Btn from "../../component/button/rounded";
import { useState } from "react";
import Username from "../../component/input/username";
import styles from "./login.module.scss";
import { cat } from "../../util/utils";

const Page: PageProps = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const elem = (
    <div className="p-3">
      <Card.Simple className="mx-auto col-10 col-sm-7 col-md-5 p-3">
        <h1 className={cat(styles.titlebar, "mx-auto")}>
          Login to <span className={styles.gotitle}>GoChat</span>
        </h1>
        <Username
          className="mt-1 w-100"
          label="username"
          value={username}
          onChange={setUsername}
        />
        <Password
          className="mt-3 w-100"
          label="password"
          value={password}
          onChange={setPassword}
        />
        <Btn.Primary className="mt-3 mx-auto py-2">login</Btn.Primary>
      </Card.Simple>
    </div>
  );
  return elem;
};

Page.layout = OneLayout;
export default Page;
