# PROGRESS LOG

## 7/4

- Added test events functionality from nearly scratch with lightning speed. Basically just my other objects as examples and tailored accordingly. Amazed at how much faster this was. Didn't really get stuck anywhere.  Things are beginning to make a lot more sense, and I feel confident that this stack could work very well for production development.

## 7/3

- today I moved all my page interfaces (handlers, services) to the domain package and now everything makes a lot more sense. Don't really understand why my interfaces were in the same place as their implementation.

## 7/2

- modified util.RenderTempl to allow adding a status code.
- continued integrating scripts into templ components and phased out form.js
- reorganized dashboard files into smaller sections
- made progress on add/delete students functionality

## 7/1

- Issue #801 with Templ [text](https://github.com/a-h/templ/issues/801) is not as severe now but still occasionally pops up. Unable to get autocomplete when working in templ files, makes development more difficult.
- Finally understand how to add javascript to templ components, by wrapping in IIFE (Immediately Invoked Function Expression). Cool stuff! Actually makes adding JS pretty easy.
- Starting to really feel like this stack could rival NextJS in terms of developer experience.
- Tried merging dev to main, resulted in merge conflict which I'm not sure how to resolve. Need to figure this out (preferred) or delete main.
- still brainstorming how I might be able to create a data structure that holds urls so that I don't have to write the same url in the component (hx-get, hx-post, etc) as well as in the handler. Seems like an unnecessary source of cognitive overhead that could easily be resolved by exposing a single source of truth about paths for data mutation and view changes.
