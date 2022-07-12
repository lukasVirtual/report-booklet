import axios from "axios";
import { useRouter } from "vue-router";

//TODO Abstraction

// export default interface loginServiceInterface {
//   login(name: string, password: string): Promise<boolean>;
//   logout(): Promise<boolean>;
//   register(name: string, password: string): Promise<boolean>;
//   checkStatus(): any;
// }

export class loginService /* implements loginServiceInterface */ {
  router = useRouter();

  static async login(name: string, password: string): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
    });

    const res = await fetch("http://127.0.0.1:5000/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: inputData,
    });
    console.log(res.body);
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

  static async register(name: string, password: string, role: string): Promise<boolean> {
    const inputData = JSON.stringify({
      name: name,
      password: password,
      role: role,
    });

    const res = await fetch("http://127.0.0.1:5000/api/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: inputData,
    });
    console.log(res.body);
    return res.status === 200 ? true : false;
  }

  static async checkStatus(): Promise<boolean> {
    const res = await axios.get("http://127.0.0.1:5000/api/statuscheck", {
      withCredentials: true,
    });
    return res.status === 200 ? true : false;
  }

  static async getUserRole(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/user", {
      withCredentials: true,
    });
    return res.data.Role;
  }

  static async getAllData(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/dataware");
    return res.data;
  }

  static async deleteItem(name: string): Promise<void> {
    const input = JSON.stringify({
       name: name,
    })
    const res = await fetch("http://127.0.0.1:5000/api/delete", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: input,
    });

    console.log(res.status)
  }
}
