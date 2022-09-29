# Exercises

Source: https://etutorials.org/Programming/Software+engineering+and+computer+games/Part+I+Software+Engineering+and+Computer+Games/Chapter+5.+Software+design+patterns/Exercises/

## Exercise 5.1: Strategy Pattern and sales chart display
Suppose that you have a cSalesData object with a drawGraph() method. Draw a UML including a few lines of code to show how you could use a Strategy Pattern to select at runtime between drawing a bar graph or a pie graph of the data.

## Exercise 5.2: Template Method pattern and opening diverse file types
When you use Visual Studio you can open various sorts of files, edit them, and then save them. Text files are shown as text and bitmap files are shown as images. We might suppose that there is a FileOpen(CString filename) method. Some of this code will act the same no matter what kind of file you open, but part of the code will act differently depending on the kind of file. Describe how this might be accomplished by using the Template Method. Draw a UML diagram and write out some very rough code, simply using made-up names for the functions you use.

## Exercise 5.3: Command pattern and a word-processor's Undo and Redo
Experiment a bit with the effect of Ctrl+Z (for Undo) and Ctrl+Y (for Redo) in your wordprocessor. What do you think might be a WordProcessorCommand in this context? What are the methods that WordProcessorCommand must implement? What kind of data structure might you need for holding the WordProcessorCommand objects? Draw a simple UML class diagram.

## Exercise 5.4: Composite pattern and building a virtual city
Say that you want to develop a cStructure class for describing doors, walls, rooms, floors of buildings, buildings, city-blocks, cities, and so on. Draw a UML showing how to do this using the Composite pattern.

## Exercise 5.5: Singleton pattern and preserving a connection
On many home computers you can to click some desktop icon to connect to the Net via a conventional modem or a broadband modem. Some email and browser software will try and make a connection even if one exists. Trying to make a connection when a connection is already open sometimes spoils the existing connection in such a way that it's impossible to reconnect without rebooting the machine. How might your operating system use the singleton pattern to avoid this problem? Draw a UML diagram.

## Exercise 5.6: Bridge pattern and the look and feel of windows
In Java it's possible to change the 'look and feel' of your windows. This includes, for instance, how you draw a window frame, draw the caption bar, draw a button, and draw a scroll bar. You can select, for instance, a Windows, a Mac, or an XWindows look and feel. Draw a UML diagram showing how this can be done by using the Bridge pattern.

## Exercise 5.7: Document-View pattern and guestbook conversation
Say that a number of clients are viewing a web page on a server. Suppose that the web page holds a guestbook that a client can update in real time by typing in characters and pressing a Send button. Draw a sequence diagram to show the steps by which Client1 can send a message to and get an answer from Client2 by using the server.