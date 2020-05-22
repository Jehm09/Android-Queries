package com.example.domainqueries.model;

public class Server {
    private String address;
    private String ssl_grade;
    private String cuntry;
    private String owner;

    public Server(String address, String ssl_grade, String cuntry, String owner) {
        this.address = address;
        this.ssl_grade = ssl_grade;
        this.cuntry = cuntry;
        this.owner = owner;
    }

    public Server(){

    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getSsl_grade() {
        return ssl_grade;
    }

    public void setSsl_grade(String ssl_grade) {
        this.ssl_grade = ssl_grade;
    }

    public String getCuntry() {
        return cuntry;
    }

    public void setCuntry(String cuntry) {
        this.cuntry = cuntry;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }
}
