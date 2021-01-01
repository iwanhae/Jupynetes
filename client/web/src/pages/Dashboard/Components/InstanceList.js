import React, { Component } from 'react'
import { Button, Card, Row, Col} from 'antd';
import Instance from './Instance.js';
import CreateInstance from '../../CreateInstance/CreateInstance';

export default class InstanceList extends Component {
    render() {
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
                    <Instance/>
                    <Instance/>
                    <Instance/>
                    <Instance/>
                    <Instance/>
                </Card>
            </>
        )
    }
}
