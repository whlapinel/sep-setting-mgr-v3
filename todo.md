# To Do

## PENDING

- Put dashboard and admin top menu buttons in a sidebar.
- Maybe allow users to enter A/B day or 4x4 though the data won't be used by the app?
- refresh token before expiration if user is active
- consolidate and simplify admin and dashboard calendar/assignment handling and services to reduce duplicate code and complexity

## COMPLETE

- implement Sign In With Google
- Calendar table heights should be consistent with each other within the week view. Maybe list every room each day instead of only listing those for which there are tests.
- admin functionality (users)
- AssignRoom handler should return more precise data instead of re-rendering entire calendar. Should render component that holds assignments for a given block and room. (need to make this component first!)
- Unassigned currently show up as overbooked (overbooked tracker is tracking assignments with room id of -1)
- optimize code for checking overbooked in admin calendar (currently checks every room for every assignment, should only do the check once and then consult a map or slice that holds roomid and boolean)
- complete edit students functionality (including room assignments, the biggest pain in the butt!!)
- display warning when room is overbooked (for each day, need counter for assignment.Event.Block, show warning for each block where the count is greater than the assignment.Room.MaxCapacity). Also need to account for student.OneOnOne; if student.OneOnOne for any assignments on that date and block, then max should be 1 instead of assignment.Room.MaxCapacity.
- admin functionality (rooms)
- edit test events functionality
- test rooms functionality
- deleting room should nullify roomID for all assignments for the room
- deleting student should delete all assignments for the student
- deleting event should delete all assignments for the event
- creating event should create assignments for every student in the event's class
- make dashboard calendar work like the admin calendar (for some reason it's not showing the assignments where room is nil, these call different repo methods which leads me to the next item)
- edit students functionality
- edit classes functionality
- change dashboard target link to dashboard/classes (add class button currently shows up in dashboard/calendar resulting in error when submitting due to no classes-table being present)
- error: GetRoomAssignments is getting passed a value of 0 for block resulting in len(assignments) always being 0
- change hx-post url in add-student-form component and url of handler
- add add student handler
- add students functionality: currently when you click add student, the form that is returned is the "add class" form. This needs to be changed to the add student form. This is because the Add Student button simply shows the hidden dialog with id of "dialog". So it opens the first dialog in the tree. Form component needs to be made so that more than one form can be present in the tree by modifying the Form (data) type and Form (templ) component, as well as the Javascript code in the Form (templ) component, so that clicking "Add Student" will show the dialog with a custom form id passed in  
- add signin check to signin page handler

## CANCELED

- Nullify should set roomid to -1 (maybe?) for sake of consistency (due to foreign constraint room_id would also need to create a room)
- dashboard: URL params should be changed to query params e.g. dashboard/classes/:class-id?students=true&test-events=true so that display state can be reflected accurately in the URL (this seemed tricky to implement so I'm putting it on the backburner for now, hopefully revisiting later)
