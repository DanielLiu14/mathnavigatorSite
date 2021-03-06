"use strict";
require("./app.styl");
import React from "react";
import ReactDOM from "react-dom";
import { withRouter } from "react-router";
import { HashRouter as Router, Route, Switch } from "react-router-dom";
import { HeaderSection } from "./header/header.js";
import { HomePage } from "./home/home.js";
import { ProgramPage } from "./programs/program.js";
import { ProgramEditPage } from "./programs/programEdit.js";
import { ClassAllPage } from "./classes/classAll.js";
import { ClassEditPage } from "./classes/classEdit.js";
import { AchievePage } from "./achieve/achieve.js";
import { AchieveEditPage } from "./achieve/achieveEdit.js";
import { AnnouncePage } from "./announce/announce.js";
import { AnnounceEditPage } from "./announce/announceEdit.js";
import { LocationPage } from "./location/location.js";
import { LocationEditPage } from "./location/locationEdit.js";
import { SemesterPage } from "./semester/semester.js";
import { SemesterEditPage } from "./semester/semesterEdit.js";

const Achieve = () => <AchievePage />;
const AchieveEdit = () => <AchieveEditPage />;
const AchieveEditMatch = ({ match }) => (
    <AchieveEditPage Id={match.params.Id} />
);
const Announce = () => <AnnouncePage />;
const AnnounceEdit = () => <AnnounceEditPage />;
const AnnounceEditMatch = ({ match }) => (
    <AnnounceEditPage announceId={match.params.announceId} />
);
const Header = () => <HeaderSection />;
const Home = () => <HomePage />;
const Programs = () => <ProgramPage />;
const ProgramEdit = () => <ProgramEditPage />;
const ProgramEditMatch = ({ match }) => (
    <ProgramEditPage programId={match.params.programId} />
);
const Location = () => <LocationPage />;
const LocationEdit = () => <LocationEditPage />;
const LocationEditMatch = ({ match }) => (
    <LocationEditPage locationId={match.params.locationId} />
);
const Semester = () => <SemesterPage />;
const SemesterEdit = () => <SemesterEditPage />;
const SemesterEditMatch = ({ match }) => (
    <SemesterEditPage semesterId={match.params.semesterId} />
);

const ClassAll = () => <ClassAllPage />;
const ClassEdit = () => <ClassEditPage />;
const ClassEditMatch = ({ match }) => (
    <ClassEditPage classId={match.params.classId} />
);

class AppContainer extends React.Component {
    render() {
        return (
            <Router>
                <AppWithRouter />
            </Router>
        );
    }
}

class App extends React.Component {
    render() {
        return (
            <div>
                <Header />
                <Switch>
                    <Route path="/" exact component={Home} />
                    <Route
                        path="/program/:programId/edit"
                        component={ProgramEditMatch}
                    />
                    <Route path="/programs/add" component={ProgramEdit} />
                    <Route path="/programs" component={Programs} />
                    <Route
                        path="/announcements/:announceId/edit"
                        component={AnnounceEditMatch}
                    />
                    <Route path="/announcements/add" component={AnnounceEdit} />
                    <Route path="/announcements" component={Announce} />
                    <Route
                        path="/achievements/:Id/edit"
                        component={AchieveEditMatch}
                    />
                    <Route path="/achievements/add" component={AchieveEdit} />
                    <Route path="/achievements" component={Achieve} />
                    <Route
                        path="/locations/:locationId/edit"
                        component={LocationEditMatch}
                    />
                    <Route path="/locations/add" component={LocationEdit} />
                    <Route path="/locations" component={Location} />
                    <Route
                        path="/semesters/:semesterId/edit"
                        component={SemesterEditMatch}
                    />
                    <Route path="/semesters/add" component={SemesterEdit} />
                    <Route path="/semesters" component={Semester} />
                    <Route
                        path="/classes/:classId/edit"
                        component={ClassEditMatch}
                    />
                    <Route path="/classes/add" component={ClassEdit} />
                    <Route path="/classes" component={ClassAll} />
                </Switch>
            </div>
        );
    }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
