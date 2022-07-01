interface loginServiceInterface {
  login(): any;
  logout(): any;
  register(): any;
  checkStatus(): any;
}

export class loginService implements loginServiceInterface {
  login() {
    throw new Error("Method not implemented.");
  }
  logout() {
    throw new Error("Method not implemented.");
  }
  register() {
    throw new Error("Method not implemented.");
  }
  checkStatus() {
    throw new Error("Method not implemented.");
  }
}
