/* -------------------------------------------------------------------------------------/
                                                                                    /
/               COPYRIGHT (c) 2019 SAMSUNG ELECTRONICS CO., LTD.                      /
/                          ALL RIGHTS RESERVED                                        /
/                                                                                     /
/   Permission is hereby granted to licensees of Samsung Electronics Co., Ltd.        /
/   products to use or abstract this computer program for the sole purpose of         /
/   implementing a product based on Samsung Electronics Co., Ltd. products.           /
/   No other rights to reproduce, use, or disseminate this computer program,          /
/   whether in part or in whole, are granted.                                         / 
/                                                                                     /
/   Samsung Electronics Co., Ltd. makes no representation or warranties with          /
/   respect to the performance of this computer program, and specifically disclaims   /
/   any responsibility for any damages, special or consequential, connected           /
/   with the use of this program.                                                     /
/                                                                                     /
/-------------------------------------------------------------------------------------/


DESCRIPTION: User Management Test File
@NAME : index.test.js
@AUTHORS: Aswin K K
@Version : 1.0 *
@REVISION HISTORY
[29/11/2019] [Aswin K K] : Prototyping..........////////////////////
*/

import React from "react";
import {
  render,
  fireEvent,
  cleanup,
  waitForElement,
  getAllByPlaceholderText,
  queryByText
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
import rootSaga from "../../../sagas/indexSaga";
import headerReducer from "../../../store/reducers/headerReducer";
import alertManagementReducer from "../../../store/reducers/alertManagementReducer";
import userManagementReducer from "../../../store/reducers/userManagementReducer";
import i18n from "../../../i18n";
import UserManagement from "./index";

jest.unmock("axios");

describe("ConfigurationSetting", () => {
  let wrapper;
  let history;
  let store;
  // let mock;
  beforeEach(() => {
    const sagaMiddleware = createSagaMiddleware();
    const rootReducers = combineReducers({
      // headerLanguageReducer,
    //   headerReducer,
      alertManagementReducer,
      userManagementReducer
    });
    const composeEnhancers =
      window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
    store = createStore(
      rootReducers,
      composeEnhancers(applyMiddleware(sagaMiddleware))
    );
    sagaMiddleware.run(rootSaga);
    const route = "/ConfigurationSetting/user";
    history = createMemoryHistory({ initialEntries: [route] });
    // mock = new MockAdapter(axios);
  });

  const renderComponent = () => {
    wrapper = render(
      <Router history={history}>
        <I18nextProvider i18n={i18n}>
          <Provider store={store}>
            <UserManagement />
          </Provider>
        </I18nextProvider>
      </Router>
    );
  };

  afterEach(cleanup);

  it('renders user management', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/').reply(200, null);
    renderComponent();
    const { getByText, asFragment } = wrapper;
    expect(getByText("User List")).toBeDefined();
  });

  it('should add a new user', () => {
    renderComponent();
    const { asFragment, getByTestId, getByText } = wrapper;
    const username = getByTestId('add-user-name');
    const password = getByTestId('add-user-password');
    const confirmPassword = getByTestId('add-user-confirm-password');
    const phno = getByTestId('add-user-phno');
    const email = getByTestId('add-user-email');
    const confirmBtn = getByText('Submit');
    fireEvent.keyDown(username, {key: 'A', code: 65, charCode: 65});
    fireEvent.keyDown(username, {key: '+', code: 43, charCode: 43});
    fireEvent.change(username, {target: { value: 'abcd'}});
    fireEvent.keyDown(password, {key: 'A', code: 65, charCode: 65});
    fireEvent.keyDown(username, {key: '+', code: 43, charCode: 43});
    fireEvent.change(password, {target: { value: 'test'}});
    fireEvent.keyDown(confirmPassword, {key: 'A', code: 65, charCode: 65});
    fireEvent.keyDown(username, {key: '+', code: 43, charCode: 43});
    fireEvent.change(confirmPassword, {target: { value: 'test'}});
    fireEvent.keyDown(email, {key: 'A', code: 65, charCode: 65});
    fireEvent.keyDown(username, {key: '+', code: 43, charCode: 43});
    fireEvent.change(email, {target: {value: 'abcd@abc.com'}});
    fireEvent.change(phno, {target: {value: '123457890'}});
    fireEvent.click(confirmBtn);
    expect(getByText('Yes')).toBeDefined();
    fireEvent.click(getByText('Yes'));
    // expect(username.value).toBe('');
  });

  it('should cancel adding a user', () => {
    renderComponent();
    const { asFragment, getByTestId, getByText } = wrapper;
    const username = getByTestId('add-user-name');
    const cancelBtn = getByText('Cancel');
    fireEvent.keyDown(username, {key: 'A', code: 65, charCode: 65});
    fireEvent.change(username, {target: { value: 'abcd'}});
    fireEvent.click(cancelBtn);
    expect(getByText('Yes')).toBeDefined();
    fireEvent.click(getByText('No'));
    // expect(username.value).toBe('');
  });

  it('should throw error when username is not present', () => {
    renderComponent();
    const { getByText } = wrapper;
    const confirmBtn = getByText('Submit');
    fireEvent.click(confirmBtn);
    expect(getByText("Please Enter a Valid Username")).toBeDefined();
  });

  it('should add list all the users', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": "admin",
          "active": true,
          "privileges": "Create, Read, Edit, Delete"
        }
      ]);
      renderComponent();
      const { getByText, asFragment } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        expect(asFragment()).toMatchSnapshot();
      });
  });

  it('should edit a user', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": "admin",
          "active": true,
          "privileges": "Create, Read, Edit, Delete"
        }
      ])
      renderComponent();
      const { getByText, asFragment, getAllByTitle, getAllByPlaceholderText } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        const editBtn = await waitForElement(() => getAllByTitle("Edit")[0]);
        fireEvent.click(editBtn);
        const phno = await waitForElement(() => getAllByPlaceholderText("+1 (702) 123-4567")[0]);
        fireEvent.change(phno, {
          target: { value: "+1 (702) 123-4578" }
        });
        const email = await waitForElement(() => getAllByPlaceholderText("Email")[0]);
        fireEvent.change(email, {
          target: { value: "test@abc.com" }
        });
        let spy = jest.spyOn(axios, "post").mockReturnValue(200);
        const saveBtn = await waitForElement(() => getAllByTitle("Save")[0]);
        fireEvent.click(saveBtn);
        await new Promise(resolve => setTimeout(resolve, 1000));
        expect(spy).toHaveBeenCalledTimes(1);
      });
  })

  it('should throw an error if the emailid is not valid', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": "admin",
          "active": true,
          "privileges": "Create, Read, Edit, Delete"
        }
      ]);
      renderComponent();
      const { getByText, asFragment, getAllByTitle, getAllByPlaceholderText } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        const editBtn = await waitForElement(() => getAllByTitle("Edit")[0]);
        fireEvent.click(editBtn);
        const phno = await waitForElement(() => getAllByPlaceholderText("+1 (702) 123-4567")[0]);
        fireEvent.change(phno, {
          target: { value: "+1 (702) 123-4578" }
        });
        const email = await waitForElement(() => getAllByPlaceholderText("Email")[0]);
        fireEvent.change(email, {
          target: { value: "test" }
        });
        let spy = jest.spyOn(axios, "post").mockReturnValue(200);
        const saveBtn = await waitForElement(() => getAllByTitle("Save")[0]);
        fireEvent.click(saveBtn);
        await new Promise(resolve => setTimeout(resolve, 1000));
        const alertText = await waitForElement(() => getByText("Please enter a valid input"));
        expect(alertText).toBeDefined();
      });
  });

  it('should delete a user', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": "admin",
          "active": true,
          "privileges": "Create, Read, Edit, Delete"
        }
      ])
      renderComponent();
      const { getByText, asFragment, getAllByTitle, getAllByPlaceholderText } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        const deleteBtn = await waitForElement(() => getAllByTitle("Delete")[0]);
        fireEvent.click(deleteBtn);
        const saveBtn = await waitForElement(() => getAllByTitle("Save")[0]);
        fireEvent.click(saveBtn);
        await new Promise(resolve => setTimeout(resolve, 1000));
        expect(asFragment()).toMatchSnapshot();
      });
  })

  it('should disable a user', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": "admin",
          "active": true,
          "privileges": "Create, Read, Edit, Delete"
        }
      ])
      renderComponent();
      const { getByText, asFragment, getAllByTitle, getAllByPlaceholderText } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        const disableBtn = await waitForElement(() => getAllByTitle("api-enable")[0]);
        let spy = jest.spyOn(axios, "post").mockReturnValue(200);
        fireEvent.click(disableBtn);
        await new Promise(resolve => setTimeout(resolve, 1000));
        fireEvent.click(disableBtn);
        expect(spy).toBeCalled();
      });
  })

  it('should not display role', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_users/')
      .reply(200, [
        {
          "_id": "abcd",
          "email": "abcd@corp.com",
          "password": "Defg",
          "phone_number": "xx",
          "role": null,
          "active": true,
          "privileges": ""
        }
      ])
      renderComponent();
      const { queryByText, getByText } = wrapper;
      await act(async () => {
        const nameElement = await waitForElement(() => getByText("abcd"));
        expect(nameElement).toBeDefined();
        expect(queryByText("Create")).toBeNull();
      });
  })
});