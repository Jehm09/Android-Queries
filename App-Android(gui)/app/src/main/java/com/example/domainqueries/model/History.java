package com.example.domainqueries.model;

public class History {
    private String []items;

    public History(String[] items) {
        this.items = items;
    }

    public History(){

    }

    public String[] getItems() {
        return items;
    }

    public void setItems(String[] items) {
        this.items = items;
    }
}
