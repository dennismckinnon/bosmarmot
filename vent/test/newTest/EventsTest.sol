pragma solidity ^0.4.20;

contract EventsTest {

    event UpdateTable(
        string indexed name,
        uint key,
        uint blocknum,
        string somestr,
        address this,
        uint instance);

    struct row {
        uint blocknum;
        string somestr;
        address thisContract;
        uint instance;
    }

    mapping(uint => row) rows;

    function editRow(uint index, string _somestr) external {
        // This is included to give a reason why the table needs to be stored
        // Strictly speaking if I just want to construct a table that I never want
        // To access it is sufficient to emit the event as that data will be stored
        // and is readable from the event log
        uint curInstance = rows[index].instance + 1;
        rows[index] = row(block.number, _somestr, this, curInstance);

        UpdateTable("EVENT_TEST", index, block.number, _somestr, this, curInstance);
    }
}