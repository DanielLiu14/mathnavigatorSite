"use strict";
import { concat, find } from "lodash";

const subLinksPrograms = [
    { id: "program-catalog", name: "Catalog", url: "/programs" },
    { id: "announcements", name: "Announcements", url: "/announcements" },
    { id: "ask-for-help", name: "Ask For Help", url: "/ask-for-help" },
];

const subLinksAchieve = [
    {
        id: "student-achieve",
        name: "Student Achievements",
        url: "/student-achievements",
    },
    {
        id: "student-webdev",
        name: "Student Web Development",
        url: "/student-webdev",
    },
    {
        id: "student-portfolios",
        name: "Student Websites",
        url: "/student-projects",
    },
];

const mainLinks = [
    {
        id: "home",
        name: "Home",
        url: "/",
    },
    {
        id: "programs",
        name: "Programs",
        url: "/programs",
        subLinks: subLinksPrograms,
    },
    {
        id: "success",
        name: "Accomplishments",
        url: "/student-achievements",
        subLinks: subLinksAchieve,
    },
    {
        id: "contact",
        name: "Contact",
        url: "/contact",
    },
];
const allLinks = concat(mainLinks, subLinksPrograms, subLinksAchieve);

export const MainLinks = mainLinks;

export function getNavById(id) {
    return find(allLinks, { id: id });
}

export function getNavByUrl(url) {
    return find(allLinks, { url: url });
}

/* not really used */
export function isPathAt(currentPath, url) {
    if (url == getNavById("home").url) {
        // return currentPath == '/'; // Use with BrowserRouter
        return currentPath == "#/";
    } else {
        return currentPath.indexOf(url) >= 0;
    }
}
