"use strict";
require("./menuWide.sass");
import React from "react";
import { Link } from "react-router-dom";
import { Navigation, isPathAt } from "../utils/links.js";
const classnames = require("classnames");

export default class MenuWide extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            links: Navigation,
        };
    }
    render() {
        const location = this.props.location;
        const onClick = this.props.onClick;
        const items = this.state.links.map((link, i) => {
            if (link.subLinks && link.subLinks.length > 0) {
                return (
                    <SubMenu
                        key={link.id}
                        currentLink={link}
                        subLinks={link.subLinks}
                        location={location}
                    />
                );
            } else {
                return (
                    <MenuLink key={link.id} link={link} location={location} />
                );
            }
        });
        return <ul className="header-menu-wide">{items}</ul>;
    }
}

class MenuLink extends React.Component {
    render() {
        const link = this.props.link;
        var currentPath = window.location.hash;
        // var currentPath = location.pathname || ""; // Use with BrowserRouter
        const linkClass = classnames({
            active: isPathAt(currentPath, link.url),
        });
        return (
            <div className="menu-link-container">
                <Link className={linkClass} to={link.url}>
                    {link.name}
                </Link>
            </div>
        );
    }
}

class SubMenu extends React.Component {
    render() {
        const currentLink = this.props.currentLink;
        const currentPath = window.location.hash;
        // var currentPath = location.pathname || ""; // Use with BrowserRouter
        const subLinks = this.props.subLinks.map((subLink, i) => {
            var linkClass = classnames({
                active: isPathAt(currentPath, subLink.url),
            });
            return (
                <li key={i}>
                    <Link className={linkClass} to={subLink.url}>
                        {subLink.name}
                    </Link>
                </li>
            );
        });
        return (
            <div className="sub-menu">
                <MenuLink link={currentLink} />
                <ul>{subLinks}</ul>
            </div>
        );
    }
}
