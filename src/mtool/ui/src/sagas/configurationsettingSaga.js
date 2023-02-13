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

import axios from 'axios';
import { call, takeEvery, put } from 'redux-saga/effects';
import * as actionTypes from '../store/actions/actionTypes';
import * as actionCreators from '../store/actions/exportActionCreators';

export function* fetchEmailList() {
  try {
    const response = yield call([axios, axios.get], '/api/v1.0/get_email_ids/', {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    });

    const result = response.data;
    const emailList = [];
    /* istanbul ignore else */
    if (result) {
      result.forEach(email => {
        emailList.push({
          ...email,
          selected: false,
          edit: false,
        });
      });
    }
    yield put(actionCreators.fetchEmailList(emailList));
  } catch (error) {
    ;
  }
}

export function* getSmtpDetails() {
  let payload = { smtpserver:'',smtpfromemail:'',smtpusername:'',smtppassword:'',isPasswordSet:false};
  try {
    const response = yield call([axios, axios.get], '/api/v1.0/get_smtp_details/', {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    });

    const result = response.data;
    const { status } = response;
    /* istanbul ignore else */
    if (status === 200) {
      payload = {
        smtpserverip: result.smtpserverip,
        smtpserverport:result.smtpserverport,
        smtpserver: `${result.smtpserverip}:${result.smtpserverport}`,
        smtpfromemail: result.smtpfromemail,
        smtpusername: result.smtpusername,
        isPasswordSet: result.isPasswordSet,
      };
    }
  } catch (error) {
    ;
  }
  finally{
    yield put(actionCreators.setSmtpServer(payload));
    yield put(actionCreators.changeSmtpServer());
  }
}

export function* deleteSmtpDetails() {
  try {
    const response = yield call([axios, axios.post], '/api/v1.0/delete_smtp_details/', {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    });
    const { status } = response;
    /* istanbul ignore else */
    if (status === 200) {
      const payload = {
        alertOpen: true,
        istypealert: true,
        alerttype: 'info',
        alerttitle: 'Delete SMTP Configuration',
        alertdescription: 'SMTP Configuration Deleted Successfully',
      };
      yield put(actionCreators.setAlertBox(payload));
      yield put(actionCreators.deleteConfiguredSmtpServer());
    }
  } catch (error) {
    const payload = {
      alertOpen: true,
      istypealert: true,
      alerttype: 'alert',
      alerttitle: 'Delete SMTP Configuration',
      alertdescription: 'Failed to Delete the SMTP Configuration'
    };
    yield put(actionCreators.setAlertBox(payload));
  }
  finally{
    yield getSmtpDetails();
  }
}

export function* updateEmail(action) {
  try {
    const response = yield call(
      [axios, axios.post],
      '/api/v1.0/update_email/',
      action.payload, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    }
    );
    const { status } = response;
     /* istanbul ignore else */
    if (status === 200) {
      const payload = {
        alertOpen: true,
        istypealert: true,
        alerttype: 'info',
        alerttitle: 'SMTP Email',
        alertdescription: 'Email List Updated Successfully',
      };
      yield put(actionCreators.setAlertBox(payload));
     
    }
  } catch (error) {
    const payload = {
      alertOpen: true,
      istypealert: true,
      alerttype: 'alert',
      alerttitle: 'SMTP Email',
      alertdescription: 'Failed to Update the Email List'
    };
    yield put(actionCreators.setAlertBox(payload));
  }
  finally {
     yield fetchEmailList();
  }
}

export function* toggleActiveStatus(action) {
  try {
    yield call(
      [axios, axios.post],
      '/api/v1.0/toggle_email_status/',
      action.payload, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    }
    );
    yield fetchEmailList();
  } catch (error) {
    ;
    // const payload = {
    //   alertOpen: true,
    //   istypealert: true,
    //   alerttype: 'alert',
    //   alerttitle: 'Toggle Active Status',
    //   alertdescription: 'Failed to perform this operation'
    // };
    // yield put(actionCreators.setAlertBox(payload));
    yield fetchEmailList();
  }
}

export function* deleteEmailIds(action) {
  try {
    if(action.payload.toggleFlag === true){
      yield toggleActiveStatus({payload: action.payload.toggleEmail})
    }
    const response = yield call(
      [axios, axios.post],
      '/api/v1.0/delete_emailids/',
      action.payload.data, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    }
    );
    const { status } = response;
     /* istanbul ignore else */
    if (status === 200) {
      const payload = {
        alertOpen: true,
        istypealert: true,
        alerttype: 'info',
        alerttitle: 'Delete Email',
        alertdescription: 'Email ID Deleted Successfully',
      };
      yield put(actionCreators.setAlertBox(payload));
      yield fetchEmailList();
    }
  } catch (error) {
    const payload = {
      alertOpen: true,
      istypealert: true,
      alerttype: 'alert',
      alerttitle: 'Delete Email',
      alertdescription: 'Failed to Delete Email ID',
    };
    yield put(actionCreators.setAlertBox(payload));
  }
}

// export function* sendEmail(action) {
//   try {
//     const response = yield call(
//       [axios, axios.post],
//       '/api/v1.0/send_email/',
//       action.payload, {
//       headers: {
//         Accept: 'application/json',
//         'Content-Type': 'application/json',
//         'x-access-token': localStorage.getItem('token'),
//       }
//     }
//     );
//     const { status } = response;
//     document.getElementsByTagName('body')[0].style.cursor = 'default';

//     if (status === 200) {
//       const payload = {
//         alertOpen: true,
//         istypealert: true,
//         alerttype: 'info',
//         alerttitle: 'Send Email',
//         alertdescription: 'Email sent successfully',
//       };
//       yield put(actionCreators.setAlertBox(payload));
//     }
//   } catch (error) {
//     document.getElementsByTagName('body')[0].style.cursor = 'default';
//     const payload = {
//       alertOpen: true,
//       istypealert: true,
//       alerttype: 'alert',
//       alerttitle: 'Send Email',
//       alertdescription: 'Email sending failed',
//     };
//     yield put(actionCreators.setAlertBox(payload));
//   }
// }

export function* testEmail(action) {
  try {
    const response = yield call(
      [axios, axios.post],
      '/api/v1.0/test_smtpserver/',
      action.payload, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'x-access-token': localStorage.getItem('token'),
      }
    }
    );
    const { status } = response;
    document.getElementsByTagName('body')[0].style.cursor = 'default';
     /* istanbul ignore else */
    if (status === 200) {
      yield put(actionCreators.changeSmtpServer());
      const payload = {
        alertOpen: true,
        istypealert: true,
        alerttype: 'info',
        alerttitle: 'Test SMTP Server',
        alertdescription: 'SMTP server is working',
      };
      yield put(actionCreators.setAlertBox(payload));
    }
  } catch (error) {
    document.getElementsByTagName('body')[0].style.cursor = 'default';
    const payload = {
      alertOpen: true,
      istypealert: true,
      alerttype: 'alert',
      alerttitle: 'Test SMTP Server',
      alertdescription: 'SMTP server is not working',
    };
    yield put(actionCreators.setAlertBox(payload));
  }
  finally{
    yield getSmtpDetails();
  }
}

// function* downloadLogs(action) {
//   try {
//     const response = yield call(
//       [axios, axios.get],
//       '/api/v1.0/download_logs', {
//       headers: {
//         'x-access-token': localStorage.getItem('token'),
//       },
//       params: action.payload,
//       responseType: 'blob'
//     }
//     );
//     const { status } = response;
//     if (status === 200) {
//       const url = window.URL.createObjectURL(new Blob([response.data]));
//       const link = document.createElement('a');
//       link.href = url;
//       link.setAttribute('download', 'log.zip');
//       // 3. Append to html page
//       document.body.appendChild(link);
//       // 4. Force download
//       link.click();
//       // 5. Clean up and remove the link
//       link.parentNode.removeChild(link);
//     }
//   } catch (error) {
//     ;
//   }
// }


export function* configurationsettingWatcher() {
  yield takeEvery(actionTypes.SAGA_FETCH_EMAIL_LIST, fetchEmailList);
  yield takeEvery(actionTypes.SAGA_UPDATE_EMAIL, updateEmail);
  yield takeEvery(actionTypes.SAGA_TOGGLE_ACTIVE_STATUS, toggleActiveStatus);
  // yield takeEvery(actionTypes.SAGA_SEND_EMAIL, sendEmail);
  yield takeEvery(actionTypes.SAGA_TEST_EMAIL, testEmail);
  yield takeEvery(actionTypes.SAGA_DELETE_EMAIL_IDS, deleteEmailIds);
  // yield takeEvery(actionTypes.SAGA_DOWNLOAD_LOGS, downloadLogs);
  yield takeEvery(actionTypes.SAGA_GET_SMTP_DETAILS, getSmtpDetails);
  yield takeEvery(actionTypes.SAGA_DELETE_SMTP_DETAILS, deleteSmtpDetails);
}
