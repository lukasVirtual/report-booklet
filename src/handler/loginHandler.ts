import axios from "axios";
import { useRouter } from "vue-router";

export default interface loginServiceInterface {
  login(name: string, password: string): Promise<boolean>;
  logout(): Promise<boolean>;
  register(name: string, password: string): Promise<boolean>;
  checkStatus(): any;
}

export class loginService implements loginServiceInterface {
  router = useRouter();
  async login(name: string, password: string): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
    });
    const res = await axios.post("http://127.0.0.1:3000/login", inputData);
    return res.status === 200 ? true : false;
  }
  async logout(): Promise<boolean> {
    const res = await axios.post("http://127.0.0.1:3000/logout");
    return res.status === 200 ? true : false;
  }
  async register(name: string, password: string): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
    });

    const res = await axios.post("http://127.0.0.1:3000/register", inputData);
    return res.status === 200 ? true : false;
  }
  checkStatus() {
    throw new Error("Method not implemented.");
  }
}
