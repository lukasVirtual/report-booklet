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
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: input,
    });

    console.log(res);
    console.log("status from saving form: ", res.status);
  }

  static async saveTextField(date: string): Promise<void> {
    const input = JSON.stringify({
      calendarDate: date,
    });

    const res = await fetch("http://127.0.0.1:5000/api/inserttext", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: input,
    });

    console.log(res);
    console.log("status from saving TextField: ", res.status);
  }

  static async getTextFieldData(date: string): Promise<any> {
    const input = JSON.stringify({
      calendarDate: date,
    });

    const res = await fetch("http://127.0.0.1:5000/api/getTextFieldData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: input,
    });

    console.log("status getting TextField data", res.status);
    return res.json();
  }

  static async writeJson(
    Id: number,
    Date: string,
    Input: string,
    Time: string,
    Rows: number
  ): Promise<void> {
    const values = JSON.stringify({
      id: Id,
      date: Date,
      input: Input,
      time: Time,
      rows: Rows,
    });
    const res = await fetch("http://127.0.0.1:5000/api/writeJson", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  static async readJson(date: string): Promise<[]> {
    const values = JSON.stringify({
      date: date,
    });
    const res = await fetch("http://127.0.0.1:5000/api/readJson", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });
    return res.json();
  }

  static async removeJson(date: string, index: number): Promise<void> {
    const values = JSON.stringify({
      index: index,
      date: date,
    });

    const res = await fetch("http://127.0.0.1:5000/api/removeJson", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  static async insertQualifications(
    qualis: {
      text: string;
      state: boolean;
    }[],
    date: string
  ): Promise<void> {
    const values = JSON.stringify({
      date: date,
      qualifications: qualis,
    });

    const res = await fetch("http://127.0.0.1:5000/api/insertQualifications", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  static async getQualifications(date: string): Promise<any> {
    const values = JSON.stringify({
      date: date,
    });

    const res = await fetch("http://127.0.0.1:5000/api/getQualifications", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    return res.json();
  }

  static async WriteStatus(date: string, status: string): Promise<void> {
    const values = JSON.stringify({
      date: date,
      status: status,
    });

    const res = await fetch("http://127.0.0.1:5000/api/writeStatusJson", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  static async ReadStatus(date: string): Promise<[]> {
    const values = JSON.stringify({
      date: date,
    });

    const res = await fetch("http://127.0.0.1:5000/api/readStatusJson", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    return res.json();
  }

  static async AssignUserToInstructor(
    instructorsName: string,
    usersName: string
  ): Promise<void> {
    const values = JSON.stringify({
      instructorsName: instructorsName,
      usersName: usersName,
    });
    const res = await fetch("http://127.0.0.1:5000/api/assigne", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  static async GetAllUserForInstructor(instructor: string): Promise<[]> {
    const values = JSON.stringify({
      instructorsName: instructor,
    });
    const res = await fetch("http://127.0.0.1:5000/api/returnUsers", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    return res.json();
  }

  static async DelteUsersFromInstructor(instructor: string): Promise<void> {
    const values = JSON.stringify({
      instructorsName: instructor,
    });

    const res = await fetch("http://127.0.0.1:5000/api/delteUserInstructor", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(`delted all users from instructor with status: ${res.status}`);
  }
}
