import axios from "axios";

export interface DataServiceInterface {
  getAllData(): Promise<any>;
  deleteItem(name: string): Promise<void>;
  saveForm(text: string, time: string): Promise<void>;
  saveTextField(date: string): Promise<void>;
  getTextFieldData(date: string): Promise<any>;
  writeJson(
    Id: number,
    Date: string,
    Input: string,
    Time: string,
    Rows: number
  ): Promise<void>;
  readJson(date: string): Promise<[]>;
  removeJson(date: string, index: number): Promise<void>;
  insertQualifications(
    qualis: {
      text: string;
      state: boolean;
    }[],
    date: string
  ): Promise<void>;
  getQualifications(date: string): Promise<any>;
  WriteStatus(date: string, status: string): Promise<void>;
  ReadStatus(date: string): Promise<[]>;
  AssignUserToInstructor(
    instructorsName: string,
    usersName: string
  ): Promise<void>;
  GetAllUserForInstructor(instructor: string): Promise<[]>;
  DelteUsersFromInstructor(instructor: string): Promise<void>;
  ExportAsPDF(): Promise<void>;
  insertCurriculum(person: string, qualis: any[]): Promise<void>;
  readCurriculum(person: string): Promise<any>;
  GetInstructorForUser(username: string): Promise<string>;
  GetJsonMonth(month: string): Promise<[]>;
  SaveSubmittedData(person: string, data: []): Promise<void>;
  RetrieveSubmittedData(person: string, month: string): Promise<[]>;
  WriteLog(message: string, instructor: string): Promise<void>;
  RetrieveLog(instructor: string): Promise<[]>;
}

let instance: DataService | undefined = undefined;

export function ProvideDataService(): DataService {
  if (!instance) instance = new DataService();
  return instance;
}

class DataService implements DataServiceInterface {
  async getAllData(): Promise<any> {
    const res = await axios.get("http://127.0.0.1:5000/api/dataware");
    return res.data;
  }

  async deleteItem(name: string): Promise<void> {
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

  async saveForm(text: string, time: string): Promise<void> {
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

  async saveTextField(date: string): Promise<void> {
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

  async getTextFieldData(date: string): Promise<any> {
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

  async writeJson(
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

  async readJson(date: string): Promise<[]> {
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

  async removeJson(date: string, index: number): Promise<void> {
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

  async insertQualifications(
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

  async getQualifications(date: string): Promise<any> {
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
    console.log(res.status);
    return res.json();
  }

  async WriteStatus(date: string, status: string): Promise<void> {
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

  async ReadStatus(date: string): Promise<[]> {
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

  async AssignUserToInstructor(
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

  async GetAllUserForInstructor(instructor: string): Promise<[]> {
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

  async DelteUsersFromInstructor(instructor: string): Promise<void> {
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

  async DelteSingleUserFromInstructor(user: string): Promise<void> {
    const values = JSON.stringify({
      usersName: user,
    });

    const res = await fetch("http://127.0.0.1:5000/api/removeUserInstructor", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  async ExportAsPDF(): Promise<void> {
    const res = await axios.get("http://127.0.0.1:5000/api/createdpdf");
    if (res.status !== 200) {
      throw new Error("Error while Creating Pdf");
    }
    const binaryString = window.atob(res.data);
    const binaryLen = binaryString.length;
    const bytes = new Uint8Array(binaryLen);
    for (let i = 0; i < binaryLen; i++) {
      const ascii = binaryString.charCodeAt(i);
      bytes[i] = ascii;
    }
    const blob = new Blob([bytes], { type: "application/pdf" });
    const href = URL.createObjectURL(blob);
    const a = Object.assign(document.createElement("a"), {
      href,
      style: "display:none",
      download: "report.pdf",
    });
    document.body.appendChild(a);
    a.click();
    URL.revokeObjectURL(href);
    a.remove();
  }

  async insertCurriculum(person: string, qualis: any[]): Promise<void> {
    const values = JSON.stringify({
      date: person,
      qualifications: qualis,
    });
    const res = await fetch("http://127.0.0.1:5000/api/insertcurriculum", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });
    console.log(res.status);
  }

  async readCurriculum(person: string): Promise<any> {
    const values = JSON.stringify({
      date: person,
    });

    const res = await fetch("http://127.0.0.1:5000/api/readcurriculum", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });
    return res.json();
  }

  async GetInstructorForUser(username: string): Promise<string> {
    const values = JSON.stringify({
      name: username,
    });

    const res = await fetch("http://127.0.0.1:5000/api/findInstructor", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    return res.json();
  }

  async GetJsonMonth(month: string): Promise<[]> {
    const values = JSON.stringify({
      date: month,
    });
    const res = await fetch("http://127.0.0.1:5000/api/readJsonMonth", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });
    return res.json();
  }

  async SaveSubmittedData(person: string, data: []): Promise<void> {
    const values = JSON.stringify({
      name: person,
      data: data,
    });

    const res = await fetch("http://127.0.0.1:5000/api/saveSubmittedData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }
  async RetrieveSubmittedData(person: string, month: string): Promise<[]> {
    const values = JSON.stringify({
      name: person,
      date: month,
    });

    const res = await fetch("http://127.0.0.1:5000/api/retrieveSubmittedData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
    return res.json();
  }

  async WriteLog(message: string, instructor: string): Promise<void> {
    const values = JSON.stringify({
      name: instructor,
      message: message,
    });

    const res = await fetch("http://127.0.0.1:5000/api/writeLog", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
  }

  async RetrieveLog(instructor: string): Promise<[]> {
    const values = JSON.stringify({
      name: instructor,
    });

    const res = await fetch("http://127.0.0.1:5000/api/retrieveLog", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "http://127.0.0.1:5000",
      },
      body: values,
    });

    console.log(res.status);
    return res.json();
  }
}
