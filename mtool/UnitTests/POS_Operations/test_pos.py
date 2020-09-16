import os

from rest.app import app
import requests_mock
import jwt
import datetime
from unittest import mock
from rest.rest_api.dagent.ibofos import start_ibofos

json_token = jwt.encode({'_id': "test", 'exp': datetime.datetime.utcnow(
) + datetime.timedelta(minutes=60)}, app.config['SECRET_KEY'])
ip = os.environ.get('DAGENT_HOST', 'localhost')

DAGENT_URL = 'http://' + ip + ':3000'

@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",
            return_value="test", autospec=True)
def test_is_ibofos_running(mock_get_current_user,**kwargs):
    kwargs["mock"].get(DAGENT_URL + '/api/dagent/v1/heartbeat',
                       json={"result": {"status": {"description": "SUCCESS", "code":0}}},
                       status_code=200)
    response = app.test_client().get(
        '/api/v1.0/get_Is_Ibof_OS_Running/',
        headers={'x-access-token': json_token})
    assert response.status_code == 200

@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",return_value="test", autospec=True)
@mock.patch("rest.app.IBOF_OS_Running.Is_Ibof_Os_Running_Flag",return_value=False, autospec=True)
def test_start_ibofos(mock_get_current_user,mock_Is_Ibof_Os_Running_Flag,**kwargs):
    kwargs["mock"].post(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "SUCCESS", "code":0}}},         
                       status_code=200)
    kwargs["mock"].get(DAGENT_URL + '/api/ibofos/v1/array',
                       status_code=200)

    response = app.test_client().get(
        '/api/v1.0/start_ibofos',
        headers={'x-access-token': json_token})

    #data = json.loads(response.get_data(as_text=True))

    assert response.status_code == 200
@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",return_value="test", autospec=True)
def test_start_ibofos_dagent_ibofos(mock_get_current_user,**kwargs):
    kwargs["mock"].post(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "SUCCESS", "code":0}}},
                       status_code=200)
    response = start_ibofos()
    assert response.status_code == 200


@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",
            return_value="test", autospec=True)
def test_start_ibofos_failure(mock_get_current_user,**kwargs):
    kwargs["mock"].post(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "FAIL", "code":20}}},
                       status_code=400)
    kwargs["mock"].get(DAGENT_URL + '/api/ibofos/v1/array',
                       status_code=200)

    response = app.test_client().get(
        '/api/v1.0/start_ibofos',
        headers={'x-access-token': json_token})

    #data = json.loads(response.get_data(as_text=True))
    assert response.status_code == 200


@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",
            return_value="test", autospec=True)
def test_stop_ibofos_already_running(mock_get_current_user,**kwargs):
    kwargs["mock"].get(DAGENT_URL + '/api/ibofos/v1/system/mount',
                       status_code=200)
    kwargs["mock"].post(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "FAIL", "code":20}}},
                       status_code=400)
    response = app.test_client().get(
        '/api/v1.0/stop_ibofos',
        headers={'x-access-token': json_token})

    #data = json.loads(response.get_data(as_text=True))
    assert response.status_code == 200



@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",
            return_value="test", autospec=True)
@mock.patch("rest.app.IBOF_OS_Running.Is_Ibof_Os_Running_Flag",
            return_value=True, autospec=True)
def test_stop_ibofos(mock_get_current_user,mock_Is_Ibof_Os_Running_Flag,**kwargs):
    kwargs["mock"].delete(DAGENT_URL + '/api/ibofos/v1/system/mount',
                       json={"result": {"status": {"description": "SUCCESS", "code":0}}},
                       status_code=200)



    kwargs["mock"].delete(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "SUCCESS", "code":0}}},
                       status_code=200)

    response = app.test_client().get(
        '/api/v1.0/stop_ibofos',
        headers={'x-access-token': json_token})

    #data = json.loads(response.get_data(as_text=True))
    assert response.status_code == 200


@requests_mock.Mocker(kw="mock")
@mock.patch("rest.app.connection_factory.get_current_user",
            return_value="test", autospec=True)
@mock.patch("rest.app.IBOF_OS_Running.Is_Ibof_Os_Running_Flag",
            return_value=True, autospec=True)
def test_stop_ibofos_failure(mock_get_current_user,mock_Is_Ibof_Os_Running_Flag,**kwargs):
    kwargs["mock"].delete(DAGENT_URL + '/api/ibofos/v1/system',
                       json={"result": {"status": {"description": "FAIL", "code":0}}},
                       status_code=400)

    response = app.test_client().get(
        '/api/v1.0/stop_ibofos',
        headers={'x-access-token': json_token})

    #data = json.loads(response.get_data(as_text=True))
    assert response.status_code == 200

