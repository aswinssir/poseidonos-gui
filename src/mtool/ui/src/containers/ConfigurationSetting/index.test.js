/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Samsung Electronics Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

import React from "react";
import {
  render,
  fireEvent,
  cleanup,
  waitForElement
} from "@testing-library/react";
import { Provider } from "react-redux";
import { act } from "react-dom/test-utils";
import { I18nextProvider } from "react-i18next";
import axios from "axios";
import "@testing-library/jest-dom/extend-expect";
import MockAdapter from "axios-mock-adapter";
import { createMemoryHistory } from "history";
import { Router } from "react-router-dom";
import { createStore, combineReducers, applyMiddleware, compose } from "redux";
import createSagaMiddleware from "redux-saga";
import rootSaga from "../../sagas/indexSaga";
import headerReducer from "../../store/reducers/headerReducer";
import configurationsettingReducer from "../../store/reducers/configurationsettingReducer";
import alertManagementReducer from "../../store/reducers/alertManagementReducer";
import ConfigurationSetting from "./index";
import i18n from "../../i18n";

jest.unmock("axios");

describe("ConfigurationSetting", () => {
  let wrapper;
  let history;
  let store;
  beforeEach(() => {
    const sagaMiddleware = createSagaMiddleware();
    const rootReducers = combineReducers({
      alertManagementReducer,
      headerReducer,
      configurationsettingReducer
    });
    const composeEnhancers =
      window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
    store = createStore(
      rootReducers,
      composeEnhancers(applyMiddleware(sagaMiddleware))
    );
    sagaMiddleware.run(rootSaga);
    const route = "/ConfigurationSetting/general";
    history = createMemoryHistory({ initialEntries: [route] });
    // mock = new MockAdapter(axios);
  });

  const renderComponent = () => {
    wrapper = render(
      <Router history={history}>
        <I18nextProvider i18n={i18n}>
          <Provider store={store}>
            {" "}
            <ConfigurationSetting />
          </Provider>
        </I18nextProvider>
      </Router>
    );
  };

  afterEach(cleanup);

  it("matches snapshot", () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);

    renderComponent();
    const { asFragment } = wrapper;
    expect(asFragment()).toMatchSnapshot();
  });
  // it("navigates to alert page", async () => {
  //   renderComponent();
  //   const { getByTestId, getByText } = wrapper;
  //   const alertTab = getByTestId("alertTab")
  //   fireEvent.click(alertTab);
  //   const generalTab = getByTestId("generalTab")
  //   fireEvent.click(generalTab);
  // });

  it("throws error on providing invalid smtp server details", async () => {
    renderComponent();
    const { getByTestId, getByText } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });
    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const mock = new MockAdapter(axios);
    mock.onPost().reply(500);
    fireEvent.click(getByTestId("applyButton"));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is not working");
    fireEvent.click(getByText("OK"));
  });

  it("deletes configured smtp server", async () => {
    const mock = new MockAdapter(axios);
    mock
    .onGet(/api\/v1.0\/get_smtp_details\/*/)
    .reply(200, 
      {
        smtpserverip: 'smtp.samsung.net',
        smtpserverport: '25'
      },
    )
    .onPost(/api\/v1.0\/test_smtpserver\/*/)
    .reply(200, {})
    .onPost(/api\/v1.0\/delete_smtp_details\/*/)
    .reply(200, {})
    .onAny()
    .reply(200, []);
    renderComponent();
    const { getByTestId,getByText } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });
    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const smtpFromEmail = getByTestId("smtpFromEmail").querySelector(
      "input"
    );
    fireEvent.change(smtpFromEmail, {
      target: { value: "your_project_email@company_xyz.com" }
    });
    const smtpUsername = getByTestId("smtpUsername").querySelector(
      "input"
    );
    fireEvent.change(smtpUsername, {
      target: { value: "username" }
    });
    const smtpPassword = getByTestId("smtpPassword").querySelector(
      "input"
    );
    fireEvent.change(smtpPassword, {
      target: { value: "abc" }
    });
    
    //const mock = new MockAdapter(axios);
    //mock.onPost().reply(200);
    fireEvent.click(getByTestId("applyButton"));
    let alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is working");
    // const readOnlyField = getByTestId("readOnlyField").querySelector("input");
    // expect(readOnlyField.value).toBe("smtp.samsung.net:25");
    fireEvent.click(getByTestId("deleteButton"));

    fireEvent.click(getByText("Yes"));
  
    alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP Configuration Deleted Successfully");
    // expect(readOnlyField.value).toBe("");
  });

  it("should delete one entry in the email list table", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      },
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();
    const { getByTestId, getByText, getAllByTitle } = wrapper;
    const deleteElement = await waitForElement(
      () => getAllByTitle("Delete")[0]
    );
    fireEvent.click(deleteElement);
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    const spy = jest.spyOn(axios, "post");
    expect(alertDescription.innerHTML).toBe(
      "Are you sure you want to Delete the email?"
    );
    mock.onPost().reply(200);
    fireEvent.click(getByText("Yes"));
    expect(spy).toHaveBeenCalledWith(
      "/api/v1.0/delete_emailids/",
      { ids: ["your_email@company_xyz.com"] },
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "x-access-token": null
        }
      }
    );
  });

  it("should delete one entry in the email list table which is inactive", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 0,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      },
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();
    const { getByTestId, getByText, getAllByTitle } = wrapper;
    const deleteElement = await waitForElement(
      () => getAllByTitle("Delete")[0]
    );
    fireEvent.click(deleteElement);
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    const spy = jest.spyOn(axios, "post");
    expect(alertDescription.innerHTML).toBe(
      "Are you sure you want to Delete the email?"
    );
    mock.onPost().reply(200);
    fireEvent.click(getByText("Yes"));
    expect(spy).toHaveBeenCalledWith(
      "/api/v1.0/delete_emailids/",
      { ids: ["your_email@company_xyz.com"] },
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "x-access-token": null
        }
      }
    );
  });

  it("configures valid smtp server details and send a test email to a user", async () => {
    const mock = new MockAdapter(axios);
    mock
    .onGet(/api\/v1.0\/get_smtp_details\/*/)
    .reply(200, 
      {
        smtpserverip: 'smtp.samsung.net',
        smtpserverport: '25'
      },
    )
    .onPost(/api\/v1.0\/test_smtpserver\/*/)
    .reply(200, {})
    .onPost(/api\/v1.0\/delete_smtp_details\/*/)
    .reply(200, {})
    .onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: true
      }
    ])
    .onAny()
    .reply(200, []);
    // const mock = new MockAdapter(axios);
    // mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
    //   {
    //     active: 1,
    //     edit: false,
    //     email: "your_email@company_xyz.com",
    //     selected: true
    //   }
    // ]);
    renderComponent();
    const { getByTestId, getByText, getByTitle } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });
    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const smtpFromEmail = getByTestId("smtpFromEmail").querySelector(
      "input"
    );
    fireEvent.change(smtpFromEmail, {
      target: { value: "your_project_email@company_xyz.com" }
    });
    const smtpUsername = getByTestId("smtpUsername").querySelector(
      "input"
    );
    fireEvent.change(smtpUsername, {
      target: { value: "username" }
    });
    const smtpPassword = getByTestId("smtpPassword").querySelector(
      "input"
    );
    fireEvent.change(smtpPassword, {
      target: { value: "abc" }
    });
    //mock.onPost().reply(200);
    fireEvent.click(getByTestId("applyButton"));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is working");
    // const readOnlyField = getByTestId("readOnlyField").querySelector("input");
    // expect(readOnlyField.value).toBe("smtp.samsung.net:25");
    fireEvent.click(getByText("OK"));
/*
    const testEmailElement = await waitForElement(() =>
      getByTitle("Test Email")
    );
    fireEvent.click(testEmailElement);
    const errorDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(errorDescription.innerHTML).toBe("Email sent successfully");
    fireEvent.click(getByTestId("alertCloseButton"));
*/
  });

  it("toggles the active status of the entry in the email list table", async () => {
    const mock = new MockAdapter(axios);
    mock
    .onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 0,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ])
    .onGet(/api\/v1.0\/get_smtp_details\/*/)
    .reply(200, 
      {
        smtpserverip: 'smtp.samsung.net',
        smtpserverport: '25'
      },
    );
    renderComponent();
    const { getByTestId, getByText, getByTitle } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });
    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const smtpFromEmail = getByTestId("smtpFromEmail").querySelector(
      "input"
    );
    fireEvent.change(smtpFromEmail, {
      target: { value: "your_project_email@company_xyz.com" }
    });
    const smtpUsername = getByTestId("smtpUsername").querySelector(
      "input"
    );
    fireEvent.change(smtpUsername, {
      target: { value: "username" }
    });
    const smtpPassword = getByTestId("smtpPassword").querySelector(
      "input"
    );
    fireEvent.change(smtpPassword, {
      target: { value: "abc" }
    });
    mock.onPost().reply(200);
    fireEvent.click(getByTestId("applyButton"));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is working");
    // const readOnlyField = getByTestId("readOnlyField").querySelector("input");
    // expect(readOnlyField.value).toBe("smtp.samsung.net:25");
    fireEvent.click(getByText("OK"));
    const toggleButton = await waitForElement(() =>
      getByTestId("toggleButton")
    );
    const spy = jest.spyOn(axios, "post");
    fireEvent.click(toggleButton);
    expect(spy).toHaveBeenCalledWith(
      "/api/v1.0/toggle_email_status/",
      { emailid: "your_email@company_xyz.com", status: true },
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "x-access-token": null
        }
      }
    );
/*
    const testEmailElement = await waitForElement(() =>
      getByTitle("Test Email")
    );
    fireEvent.click(testEmailElement);
    const errorDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(errorDescription.innerHTML).toBe(
      "Please select an email id to send"
    );
*/
  });
/*
  it("throws error while trying to send an email when smtp server is not configured", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();
    const { getByTestId, getByTitle } = wrapper;
    const testEmailElement = await waitForElement(() =>
      getByTitle("Test Email")
    );
    fireEvent.click(testEmailElement);
    const errorDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(errorDescription.innerHTML).toBe("Please configure smtp server");
  });
*/
  it("throws an error if a duplicate entry is added in the email list", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();

    const {
      asFragment,
      getByTitle,
      getByPlaceholderText,
      getByTestId
    } = wrapper;
    const addElement = getByTitle("Add");
    fireEvent.click(addElement);
    await act(async () => {
      const saveElement = await waitForElement(() => getByTitle("Save"));
      expect(asFragment()).toMatchSnapshot();
      const inputNode = getByPlaceholderText("Email ID");
      fireEvent.change(inputNode, {
        target: { value: "your_email@company_xyz.com" }
      });
      fireEvent.click(saveElement);

      const alertDescription = await waitForElement(() =>
        getByTestId("alertDescription")
      );
      expect(alertDescription.innerHTML).toBe("This email id exists");
    });
  });

  it("throws an error if a blank entry is added to the email list", async () => {
    renderComponent();
    const { asFragment, getByTitle, getByTestId } = wrapper;
    const addElement = getByTitle("Add");
    fireEvent.click(addElement);
    // await act(async () => {
    const saveElement = await waitForElement(() => getByTitle("Save"));
    expect(asFragment()).toMatchSnapshot();
    fireEvent.click(saveElement);
    await new Promise(resolve => setTimeout(resolve, 1000));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("Please enter a valid email id");
    // });
  });

  it("edits an email entry in the email list", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();
    const { asFragment, getByTitle, getByPlaceholderText } = wrapper;

    await act(async () => {
      const editElement = await waitForElement(() => getByTitle("Edit"));
      fireEvent.click(editElement);
      expect(asFragment()).toMatchSnapshot();
      const saveElement = await waitForElement(() => getByTitle("Save"));
      const inputNode = getByPlaceholderText("Email ID");
      fireEvent.change(inputNode, {
        target: { value: "your_email@company_xyz.com" }
      });
      const spy = jest.spyOn(axios, "post");
      fireEvent.click(saveElement);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toHaveBeenCalledWith(
        "/api/v1.0/update_email/",
        {
          active: 1,
          edit: false,
          email: "your_email@company_xyz.com",
          oldid: "your_email@company_xyz.com",
          selected: false
        },
        {
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            "x-access-token": null
          }
        }
      );
    });
  });




  

  it("should render button on resize", () => {
    // Change the viewport to 500px.
    global.innerWidth = 500;

    // Trigger the window resize event.
    global.dispatchEvent(new Event("resize"));

    renderComponent();
    const { getByTestId } = wrapper;
    expect(getByTestId("sidebar-toggle")).toBeDefined();
    fireEvent.click(getByTestId("sidebar-toggle"));
    expect(getByTestId("help-link")).toHaveTextContent("Help");
  });

  it("adds a new entry to the email list", async () => {
    renderComponent();
    const { asFragment, getByTitle, getByPlaceholderText } = wrapper;
    const addElement = getByTitle("Add");
    fireEvent.click(addElement);
    await act(async () => {
      const saveElement = await waitForElement(() => getByTitle("Save"));
      expect(asFragment()).toMatchSnapshot();
      const inputNode = getByPlaceholderText("Email ID");
      fireEvent.change(inputNode, {
        target: { value: "your_email@company_xyz.com" }
      });
      const spy = jest.spyOn(axios, "post");
      const mock = new MockAdapter(axios);
      mock.onPost("/api/v1.0/update_email/").reply(200);
      fireEvent.click(saveElement);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toHaveBeenCalledWith(
        "/api/v1.0/update_email/", {
          active: 1,
          edit: false,
          email: "your_email@company_xyz.com",
          oldid: "your_email@company_xyz.com",
          selected: false
        },
        {
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            "x-access-token": null
          }
        }
      );
    });
  });

  it("toggle api throws error if failed", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 0,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ])
    .onGet(/api\/v1.0\/get_smtp_details\/*/)
    .reply(200, 
      {
        smtpserverip: 'smtp.samsung.net',
        smtpserverport: '25'
      },
    );
    renderComponent();
    const { getByTestId, getByText } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });
    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const smtpFromEmail = getByTestId("smtpFromEmail").querySelector(
      "input"
    );
    fireEvent.change(smtpFromEmail, {
      target: { value: "your_project_email@company_xyz.com" }
    });
    const smtpUsername = getByTestId("smtpUsername").querySelector(
      "input"
    );
    fireEvent.change(smtpUsername, {
      target: { value: "username" }
    });
    const smtpPassword = getByTestId("smtpPassword").querySelector(
      "input"
    );
    fireEvent.change(smtpPassword, {
      target: { value: "abc" }
    });
    mock.onPost().reply(200);
    fireEvent.click(getByTestId("applyButton"));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is working");
    // const readOnlyField = getByTestId("readOnlyField").querySelector("input");
    // expect(readOnlyField.value).toBe("smtp.samsung.net:25");
    fireEvent.click(getByText("OK"));
    const toggleButton = await waitForElement(() =>
      getByTestId("toggleButton")
    );
    mock.onPost().reply(500);
    const spy = jest.spyOn(axios, "post");
    fireEvent.click(toggleButton);
    expect(spy).toHaveBeenCalledWith(
      "/api/v1.0/toggle_email_status/",
      { emailid: "your_email@company_xyz.com", status: true },
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "x-access-token": null
        }
      }
    );
  });

  it("delete email api throws error", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      },
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ]);
    renderComponent();
    const { getByTestId, getByText, getAllByTitle } = wrapper;
    const deleteElement = await waitForElement(
      () => getAllByTitle("Delete")[0]
    );
    fireEvent.click(deleteElement);
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    const spy = jest.spyOn(axios, "post");
    expect(alertDescription.innerHTML).toBe(
      "Are you sure you want to Delete the email?"
    );
    mock.onPost().reply(500);
    fireEvent.click(getByText("Yes"));
    expect(spy).toHaveBeenCalledWith(
      "/api/v1.0/delete_emailids/",
      { ids: ["your_email@company_xyz.com"] },
      {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "x-access-token": null
        }
      }
    );
    const errorDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(errorDescription.innerHTML).toBe("Failed to Delete Email ID");
    fireEvent.click(getByTestId("alertCloseButton"));
  });

  it("send email api fails", async () => {
    const mock = new MockAdapter(axios);
    mock.onGet("/api/v1.0/get_email_ids/").reply(200, [
      {
        active: 1,
        edit: false,
        email: "your_email@company_xyz.com",
        selected: false
      }
    ])
    .onGet(/api\/v1.0\/get_smtp_details\/*/)
    .reply(200, 
      {
        smtpserverip: 'smtp.samsung.net',
        smtpserverport: '25'
      },
    );
    renderComponent();
    const { getByTestId, getByText, getByTitle } = wrapper;
    const smtpServerField = getByTestId("smtpServerField").querySelector(
      "input"
    );
    fireEvent.change(smtpServerField, {
      target: { value: "smtp.samsung.net:25" }
    });

    expect(smtpServerField.value).toBe("smtp.samsung.net:25");
    const smtpFromEmail = getByTestId("smtpFromEmail").querySelector(
      "input"
    );
    fireEvent.change(smtpFromEmail, {
      target: { value: "your_project_email@company_xyz.com" }
    });
    const smtpUsername = getByTestId("smtpUsername").querySelector(
      "input"
    );
    fireEvent.change(smtpUsername, {
      target: { value: "username" }
    });
    const smtpPassword = getByTestId("smtpPassword").querySelector(
      "input"
    );
    fireEvent.change(smtpPassword, {
      target: { value: "abc" }
    });
    mock.onPost().reply(200);
    fireEvent.click(getByTestId("applyButton"));
    const alertDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(alertDescription.innerHTML).toBe("SMTP server is working");
    // const readOnlyField = getByTestId("readOnlyField").querySelector("input");
    // expect(readOnlyField.value).toBe("smtp.samsung.net:25");
    fireEvent.click(getByText("OK"));
/*
    const testEmailElement = await waitForElement(() =>
      getByTitle("Test Email")
    );
    mock.onPost().reply(500);
    fireEvent.click(testEmailElement);
    const errorDescription = await waitForElement(() =>
      getByTestId("alertDescription")
    );
    expect(errorDescription.innerHTML).toBe("Email sending failed");
*/
  });
  // it("should download the logs", async () => {
  //   renderComponent();
  //   const { getByTestId } = wrapper;
  //   fireEvent.click(getByTestId("downloadLogsBtn"));
  // });
  // it("should set ibofos time interval", async () => {
    
  //   renderComponent();
  //   const { getByTestId, getByText } = wrapper;
  //   const ibofosTimeIntervalField = getByTestId("ibofosSettingTextField").querySelector(
  //     "input"
  //   );
  //   fireEvent.change(ibofosTimeIntervalField, {
  //     target: { value: "4" }
  //   });
  //   fireEvent.click(getByTestId("setTimeIntervalButton"));
  //   fireEvent.change(ibofosTimeIntervalField, {
  //     target: { value: "-1" }
  //   });
  //   fireEvent.click(getByTestId("setTimeIntervalButton"));
  //   const okBtn = getByText("OK");
  //   expect(okBtn).toBeDefined();
  //   fireEvent.click(okBtn);
  // });

  // it("should delete ibofos time interval", async () => {
  //   renderComponent();
  //   const { getByTestId } = wrapper;
  //   fireEvent.click(getByTestId("deleteTimeIntervalButton"));
  // });

  //Disabling for PoC1

  // it("should switch tabs", async () => {
  //   renderComponent();
  //   const { getByText } = wrapper;
  //   fireEvent.click(getByText("Alert"));
  //   fireEvent.click(getByText("General"));
  // });
});
