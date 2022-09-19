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
 *     * Neither the name of Intel Corporation nor the names of its
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

import React, { useState } from 'react';
import { Button, FormControl, FormControlLabel, TextField, withStyles, Checkbox } from '@material-ui/core';
import { connect } from 'react-redux';
import * as actionTypes from '../../store/actions/actionTypes';

const styles = (theme) => ({
    createSubsystemForm: {
        display: 'flex',
        flexDirection: 'column'
    },
    formItem: {
        width: '48%',
        margin: '1%',
        height: 48,
        justifyContent: 'flex-end'
    },
    submitBtn: {
        margin: theme.spacing(1)
    }
});

const defaultSubsystem = {
    "name": "",
    "allowAnyHost": false,
    "mn": "",
    "sn": "",
    "maxNamespaces": 256
};

const CreateSubsystem = (props) => {
    const [subsystem, setSubsystem] = useState(defaultSubsystem);

    const onChange = (event) => {
        const { name, value } = event.target;
        setSubsystem({
            ...subsystem,
            [name]: value
        })
    }

    const onCheckboxChange = (event) => {
        const { name, checked } = event.target;
        setSubsystem({
            ...subsystem,
            [name]: checked
        })
    }


    const { classes } = props;
    return (
        <form className={classes.createDiskForm}>
            <FormControl
                className={classes.formItem}
            >
                <TextField
                    label="SUBNQN"
                    value={subsystem.name}
                    name="name"
                    inputProps={{
                        "data-testid": "subsystemName"
                    }}
                    onChange={onChange}
                />
            </FormControl>
            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="Serial No"
                    value={subsystem.sn}
                    name="sn"
                    inputProps={{
                        "data-testid": "subsystemSN"
                    }}
                    onChange={onChange}
                />
            </FormControl>
            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="Model No"
                    value={subsystem.mn}
                    name="mn"
                    inputProps={{
                        "data-testid": "subsystemMN"
                    }}
                    onChange={onChange}
                />
            </FormControl>

            <FormControl
                className={classes.formItem}
                size="small"
            >
                <TextField
                    label="Maximum Namespaces"
                    value={subsystem.maxNamespaces}
                    name="maxNamespaces"
                    type="number"
                    inputProps={{
                        "data-testid": "subsystemMaxNS"
                    }}
                    onChange={onChange}
                />
            </FormControl>
            <FormControl
                className={classes.formItem}
            >
                <FormControlLabel
                    label="Allow Any Host?"
                    control={(
                        <Checkbox
                            name="allowAnyHost"
                            value={subsystem.allowAnyHost}
                            inputProps={{
                                id: "allowAnyHost",
                                "data-testid": "allowAnyHost"
                            }}
                            onChange={onCheckboxChange}
                        />
                    )}
                />
            </FormControl>
            <Button
                className={classes.submitBtn}
                variant="contained"
                color="primary"
                data-testid="subsystemCreate"
                onClick={() => props.createSubsystem(subsystem)}
            >
                Create Subsystem
            </Button>
        </form>
    );
}

const mapStateToProps = () => ({});

const mapDispatchToProps = (dispatch) => ({
    Create_Disk: (payload) => dispatch({ type: actionTypes.SAGA_CREATE_SUBSYSTEM, payload })
});

export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(CreateSubsystem));
