import React, { FC } from 'react'
import {  Card, Row, Col} from 'antd';
import CreateInstance from '../../CreateInstance/CreateInstance';
import appState from '../../../AppState';
import { ServerObject } from '../../../model/ServerObject';
import InstanceContainer from '../Container/InstanceContainer';
import { useObserver } from "mobx-react";

 const InstanceList: FC<Object> = () =>  {
    return useObserver(() => {
        return (
            <>
                <div id="test">
                    <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                </div>
                <div id="test">
                    <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                </div>
                <Row gutter={16}>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                    </Col>
                    <Col className="gutter-row" span={6}>
                        <CreateInstance></CreateInstance>
                    </Col>
                </Row>
                <Card title="내 인스턴스 목록">
                    {appState.servers.map(function(item:ServerObject) {
                        return <InstanceContainer serverData={item}></InstanceContainer>;
                    })}
                </Card>
            </>
        )
    });
}

export default InstanceList;

