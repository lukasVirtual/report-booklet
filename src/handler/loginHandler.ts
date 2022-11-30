import axios from "axios";
import { useRouter } from "vue-router";

// export interface LoginServiceInterface {
//   login(name: string, password: string): Promise<boolean>;
//   logout(): Promise<boolean>;
//   register(name: string, password: string, role: string): Promise<boolean>;
//   checkStatus(): any;
//   getUserRole(): Promise<any>;
//   getUser(): Promise<any>;
//   getUserID(): Promise<any>;
// }
//
// let instance: LoginService | undefined = undefined;
//
// export function ProvideLoginService(): LoginService {
//   if (!instance) instance = new LoginService();
//   return instance;
// }

export class loginService {
  router = useRouter();

  static async login(name: string, password: string): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
    });

    const res = await fetch("http://127.0.0.1:5000/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      credentials: "include",
      body: inputData,
    });
    console.log(await res.json());
    return res.status === 200 ? true : false;
  }

  static async logout(): Promise<boolean> {
    const res = await fetch("http://127.0.0.1:5000/api/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    return res.status === 200 ? true : false;
  }

  static async register(
    name: string,
    password: string,
    role: string
  ): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
      role: role,
    });

    const res = await fetch("http://127.0.0.1:5000/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: inputData,
    });
    console.log(res.status);
    return res.status === 200 ? true : false;
  }

  static async checkStatus(): Promise<boolean> {
    const res = await axios.get("http://127.0.0.1:5000/api/statuscheck", {
      withCredentials: true,
    });
    return res.status === 200 ? true : false;
  }

  /* TODO remove duplication */

  static async getUserRole(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/user", {
      withCredentials: true,
    });
    return res.data.Role;
  }
  static async getUser(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/user", {
      withCredentials: true,
    });
    return res.data.Name;
  }
  static async getUserID(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/user", {
      withCredentials: true,
    });
    return res.data.ID;
  }
}
