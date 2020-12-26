import React from "react";
// react plugin for creating charts
// @material-ui/core
import { makeStyles } from "@material-ui/core/styles";
// @material-ui/icons
import Update from "@material-ui/icons/Update";
// core components
import GridItem from "components/Grid/GridItem.js";
import GridContainer from "components/Grid/GridContainer.js";
import Card from "components/Card/Card.js";
import CardHeader from "components/Card/CardHeader.js";
import CardBody from "components/Card/CardBody.js";
import CardFooter from "components/Card/CardFooter.js";
import Container from '@material-ui/core/Container';

import Box from "@material-ui/core/Box"
import Typography from '@material-ui/core/Typography';

import Button from "components/CustomButtons/Button.js";
import styles from "assets/jss/material-dashboard-react/views/dashboardStyle.js";

const useStyles = makeStyles(styles);

export default function Instane() {

  const classes = useStyles();

  const statusStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#F00",
    borderRadius: "50%",
    display: "inline-block",
  }

  const deleteButtonStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#992222",
  }

  const accessButtonStyle={
    height: "50px",
    width: "50px",
    backgroundColor: "#009955",
  }

    return <Card>
        <Box color="warning">
        <CardHeader color="success" stats icon>
            <Button type="button" color="info"><h3>사과서버</h3></Button>
        </CardHeader>
        </Box>
        <Container>
        <Typography component="div" style={{ backgroundColor: '#cfe8fc', height: '22vh' }}>
            <CardBody>
            <p>언제 생성</p>
            <p>몇일째 구동중</p>
            <p>각종 잡다한 세부정보</p>
            <p>구동중인 이미지</p>
            </CardBody>
        </Typography>
        <GridContainer>
            <GridItem xs={12} sm={6} md={5}>
            <span class="dot" style={statusStyle}></span>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
            <Button type="button" style={deleteButtonStyle}>
                삭제
            </Button>
            </GridItem>
            <GridItem xs={12} sm={6} md={3}>
            <Button type="button" style={accessButtonStyle}>
                접속
            </Button>
            </GridItem>
        </GridContainer>
        </Container>
        <CardFooter stats>
        <div className={classes.stats}>
            <Update />
            Just Updated
        </div>
        </CardFooter>
    </Card>;
}