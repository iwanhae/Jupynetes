import React from "react";
// react plugin for creating charts
// @material-ui/core
import { makeStyles } from "@material-ui/core/styles";
// @material-ui/icons
import Update from "@material-ui/icons/Update";
// core components
import Card from "components/Card/Card.js";
import CardHeader from "components/Card/CardHeader.js";
import CardFooter from "components/Card/CardFooter.js";
import Container from '@material-ui/core/Container';

import Box from "@material-ui/core/Box"

import Button from "components/CustomButtons/Button.js";
import styles from "assets/jss/material-dashboard-react/views/dashboardStyle.js";
import Content from "./Content.js";
import Panel from "./Panel.js";

const useStyles = makeStyles(styles);

export default function Instane(props) {

  const classes = useStyles();
    return <Card>
        <Box color="warning">
        <CardHeader color="success" stats icon>
            <Button type="button" color="info"><h3>{props.children}</h3></Button>
        </CardHeader>
        </Box>
        <Container>
        <Content></Content>
        <Panel></Panel>
        </Container>
        <CardFooter stats>
        <div className={classes.stats}>
            <Update />
            Just Updated
        </div>
        </CardFooter>
    </Card>;
}