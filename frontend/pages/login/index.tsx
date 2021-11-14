import Card from "../../component/cards/rounded";
import Password from "../../component/input/password";
import OneLayout from "../../component/layout/oneLayout";
import PageProps from "../../definition/pageProps";
import { useState } from "react";
import Username from "../../component/input/username";

const Page: PageProps = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const elem = (
    <div>
      <h1>Login to GoChat</h1>
      <Card.Simple className="col-10 col-sm-7 col-md-5">
        <Username
          className="mt-3 w-100"
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
        
      </Card.Simple>
    </div>
  );
  return elem;
};

Page.layout = OneLayout;
export default Page;
