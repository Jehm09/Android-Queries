package com.example.domainqueries.model;

public class Domain {
    private Server []servers;
    private boolean servers_changed;
    private String ssl_grade;
    private String previous_ssl_grade;
    private String logo;
    private String title;
    private boolean is_down;

    public Domain(Server[] servers, boolean servers_changed, String ssl_grade, String previous_ssl_grade, String logo, String title, boolean is_down) {
        this.servers = servers;
        this.servers_changed = servers_changed;
        this.ssl_grade = ssl_grade;
        this.previous_ssl_grade = previous_ssl_grade;
        this.logo = logo;
        this.title = title;
        this.is_down = is_down;
    }

    public Domain() {

    }

    public Server[] getServers() {
        return servers;
    }

    public void setServers(Server[] servers) {
        this.servers = servers;
    }

    public boolean isServers_changed() {
        return servers_changed;
    }

    public void setServers_changed(boolean servers_changed) {
        this.servers_changed = servers_changed;
    }

    public String getSsl_grade() {
        return ssl_grade;
    }

    public void setSsl_grade(String ssl_grade) {
        this.ssl_grade = ssl_grade;
    }

    public String getPrevious_ssl_grade() {
        return previous_ssl_grade;
    }

    public void setPrevious_ssl_grade(String previous_ssl_grade) {
        this.previous_ssl_grade = previous_ssl_grade;
    }

    public String getLogo() {
        return logo;
    }

    public void setLogo(String logo) {
        this.logo = logo;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public boolean isIs_down() {
        return is_down;
    }

    public void setIs_down(boolean is_down) {
        this.is_down = is_down;
    }
}
