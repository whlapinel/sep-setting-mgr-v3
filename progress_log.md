# PROGRESS LOG

## 7/26

- Added side nav components to dashboard and admin pages
- I'm not sure if I'm on the right track now, but I'm experimenting with how I call templ components. I've been uncomfortable with using struct literals for component properties (how do I know what's required and what's optional?) so starting with the new dialog component (so simple that it's probably a bad example for this purpose, but just wanted to see how it would work) I'm creating a constructor for the dialogProps struct e.g. Dialog() and this struct has a method Templify(). So I tried using this in the templ component, but method calls don't seem to enjoy syntax highlighting.  So I created an interface Templifer with the Templify method as its requirement, and a function that accepts a Templifier, and then I call that function and pass in the constructor call. It's quite convoluted, and I'm not sure if I'm gaining much here (other than reinforcing knowledge about these things) and need to think more about whether to apply this pattern elsewhere.

## 7/24

- Finally implemented Sign In With Google, removed password fields and added first and last name.
- The docs for Sign In with Google were surprisingly sparse and confusing. Could have done this a lot faster with better docs.

## 7/23

- Continue styling and formatting calendar tables

## 7/22

- Continue styling and componentizing calendar tables
- Show warning on overbooked rooms or unassigned students
- Simplify week view with more detailed day view
- Create separate handler for dashboard calendar details

## 7/21

- Edit Users functionality today. Also began styling tables.

## 7/20

- Spent all morning re-working calendar. Using a nested map, which seemed overwhelmingly complex at first and took me a long time to get comfortable working with it. But this is a lot more efficient than what I had before.

## 7/19

- Removed assign buttons from dashboard calendar (at this point I want assign rooms to be admin privilege, in keeping with current practices)
- Fixed edit rooms form

## 7/18

- Spent most of this morning working on assign rooms functionality from the /admin/calendar page.

## 7/17

- Struggled a little bit with how to display assignments properly, all morning was spent reworking things to make it simpler and more performant. Now calendars are only passed assignments instead of both test events and assignments.
- Major change in how calendar data is retrieved: when students are created, an assignment is created for every test event in the student's class, with the room id set to null. This meant learning about sql.NullInt64 and sql.NullString and how to use them.
- Still need to do the same for test events (new assignment should be created for every student in the test event's class).
- This means I'll also need to implement auto-delete and auto-update accordingly, but I would need to do that anyway, and it will be way simpler than when I was trying to auto-assign to rooms *shudder*.

## 7/16

- I've been spending a lot of time reorganizing the project to match more closely with Go project structure conventions. I may still be deviating a bit but I feel this is closer. It also feels more sensible.
- It suddenly dawned on me today that I've got myself into a real pickle trying to auto-assign students to rooms. It adds a ton of complexity and I could put that feature in later. It would also be a lot safer to have auto-assign as a button-triggered, on-demand feature available per-date only in the admin dashboard, rather than triggering it every time a student, test event, or room change happens. It would also be better to allow the admin to confirm or modify auto-assignments before they're saved. The whole thing has suddenly become a nightmare, so I think I'm going to remove all current auto-assignment methods. That will be its own small nightmare but not nearly as bad. And then later I will implement auto-assign in the admin panel.

## 7/12

- Finally took the time to research echo capabilities and lo and behold, I found the solution to the thing that's been bugging me since I started this project. It would have saved me a lot of trouble if I did that sooner, but better late than never!  I can pass into my templates the echo instance and generate URIs using handlers or named routes. I can also just pass in paths for those without dynamic parameters. The parameters for my templates will get a bit lengthy so I think I may want to define them in a struct.
- Update: I couldn't get the echo.URI() method to work passing in my method handlers, but tests using echo.Reverse() using named routes. Very happy with this breakthrough.

## 7/11

- Worked on editing functionality. Had to rework the forms but I'm happy now.  The only thing I don't like is that I currently remove the form with JS upon clicking cancel, and that puts a little warning in the console log. But probably not really an issue unless there's something with accessibility.
- Still need to test edit students functionality but edit classes works.
- Haven't started with editing test event functionality yet.
- Overall, making great progress and learning a lot.
- Still spending more time fussing with route paths than it seems I should, would love to have a better system for auto-syncing up the hx-targets with the handlers so I'm not constantly having to remember (and forgetting).

## 7/9

- Over the last several days I've been slogging through the room assignments functionality, but have made good, albeit slow, progress.
- When students or test events are created, a room assignment is created, according to the priority of the room. If the rooms are all full, then a message is sent to the user.
- Especially excited to have implemented a client-side messaging system using HX-Trigger response headers. Beautiful system for triggering client-side events, very excited to use this.

## 7/5

- Went around in circles today. Trying to figure out the most complex part of this app, which is how to assign students to rooms automatically. Thought I could do it without creating a repository or persisting assignments and just provide assignments when display is requested. But that started to feel like it wouldn't work.

## 7/4

- Added test events functionality from nearly scratch with lightning speed. Basically just my other objects as examples and tailored accordingly. Amazed at how much faster this was. Didn't really get stuck anywhere.  Things are beginning to make a lot more sense, and I feel confident that this stack could work very well for production development.
- not in my project, but as a side note, the Templ repo issue with missing completions and recurring "Request textDocument/codeAction failed" error is resolved!!  This will be a huge help.

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

## 6/28

After completing summer 1st term I resumed working on this project in earnest. I'm having an issue with the templ-vscode extension that's really bugging me. I went to the Gopher slack and posted a question about it in the #templ channel, and got a nice response from the developer "a-h" himself, who indicated that this issue has been identified and is being addressed.  Things are continuing to work alright, aside from the annoying message that keeps popping up "Request textDocument/codeAction failed." I am not sure if there is any impact to the development process, and the code runs correctly.

