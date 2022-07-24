import axios from "axios";

export class dataService {
  static async getAllData(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/dataware");
    return res.data;
  }

  static async deleteItem(name: string): Promise<void> {
    const input = JSON.stringify({
      name: name,
    });
    const res = await fetch("http://127.0.0.1:5000/api/delete", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: input,
    });
    console.log(res.status);
  }

  static async saveForm(text: string, time: string): Promise<void> {
    const input = JSON.stringify({
      input: text,
      timeStamp: time,
    });
    const res = await fetch("http://127.0.0.1:5000/api/insertinput", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: input,
    });

    console.log(res);
    console.log("status from saving form: ", res.status);
  }
}
